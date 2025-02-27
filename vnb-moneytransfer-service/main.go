package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	"github.com/IBM/sarama"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natefinch/lumberjack"
	"github.com/rakyll/statik/fs"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tunvx/simplebank/common/logger"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	_ "github.com/tunvx/simplebank/grpc/doc/moneytransfer/statik"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	"github.com/tunvx/simplebank/moneytransfersrv/cache"
	"github.com/tunvx/simplebank/moneytransfersrv/gapi"
	worker "github.com/tunvx/simplebank/moneytransfersrv/worker/kafka"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	// Load configuration from the environment or config file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Set up lumberjack for log rotation
	logFile := &lumberjack.Logger{
		Filename:   "/var/log/service-transfermoney.log",
		MaxSize:    10,   // Maximum size in MB before rotation
		MaxBackups: 3,    // Keep at most 3 old backups
		MaxAge:     30,   // Max age in days to retain old log files
		Compress:   true, // Compress rotated log files
	}

	// Set logger based on environment
	if config.Environment == "development" {
		// In development mode, output to console in a readable format
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	} else if config.Environment == "production" {
		// In production, use a multi-writer to log to both stdout and a log file
		multi := io.MultiWriter(os.Stdout, logFile)
		log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	} else {
		// Default fallback if environment is not recognized
		log.Fatal().Msg("Unknown environment: " + config.Environment)
	}
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// STORE DB: Establish SQL Store for all shards
	stores, err := establishShardedSQLStore(config.ListDBSourceCoreDB, config.NumCoreDBShard)
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: Failed to initialize stores for shards")
	}
	defer func() {
		for _, store := range stores {
			if sqlStore, ok := store.(*db.SQLStore); ok {
				sqlStore.Close()
			}
		}
	}()

	// CACHE REDIS: Redis connection pool configuration
	redisOpt := redis.Options{
		Addr:            config.InternalRedisAddress,
		MaxActiveConns:  120,
		ConnMaxIdleTime: time.Minute * 5,
	}
	cache := cache.NewRedisCache(&redisOpt)

	// KAFKA PRODUCER: Kafka producer configuration
	brokers := []string{config.InternalKafkaAddress}
	topics := []string{worker.TopicInternalTransferMoney}

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.ClientID = "moneytransfer service"
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll // Ensure message is written to all replicas
	kafkaConfig.Producer.Retry.Max = 0                    // Retry sending if error
	kafkaConfig.Producer.Return.Successes = true          // Catch success event
	kafkaConfig.Producer.Return.Errors = true             // Catch sending errors

	taskProducer, err := worker.NewKafkaTaskProducer(brokers, topics, kafkaConfig)
	if err != nil {
		log.Fatal().Err(err).Msgf("MoneyTransfer Service: Failed to create taskProducer: %s", err)
	}

	waitGroup, ctx := errgroup.WithContext(ctx)
	runGrpcServer(ctx, waitGroup, config, stores, cache, taskProducer)
	runGatewayServer(ctx, waitGroup, config, stores, cache, taskProducer)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: Error from wait group")
	}

	// log.Info().Msg("add trigger to run github action")
}

func establishShardedSQLStore(listShardDatabaseURL []string, numShard int) ([]db.Store, error) {
	// Check if the length of listShardDatabaseURL matches numShard
	if len(listShardDatabaseURL) != numShard {
		log.Fatal().Msgf("MoneyTransfer Service: The length of listShardDatabaseURL ( %d ) does not match numShard ( %d )", len(listShardDatabaseURL), numShard)
	}

	var stores []db.Store

	// Loop over each shard and apply migration
	for shardID := 1; shardID <= numShard; shardID++ {
		// Access the shard database URL based on shardID
		shardDatabaseURL := listShardDatabaseURL[shardID-1]
		// *************************************************************
		connPool, err := pgxpool.New(context.Background(), shardDatabaseURL)
		if err != nil {
			log.Error().Err(err).Msgf("cannot connect to shard ( %d ), error: %v", shardID, err)
			continue
		} else {
			log.Info().Msgf("successfully created pool connection to shard ( %d )", shardID)
		}
		// defer connPool.Close()

		// Create a new store to interact with the database
		store := db.NewStore(connPool)
		stores = append(stores, store)
	}

	if len(stores) == 0 {
		return nil, fmt.Errorf("MoneyTransfer Service: Failed to initialize any database shard")
	}

	return stores, nil
}

// runGrpcServer starts the gRPC server for handling core banking services
func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	stores []db.Store,
	cache cache.Cache,
	taskProducer worker.TaskProducer,
) {
	// Create a new transaction service
	tranService, err := gapi.NewService(config, stores, cache, taskProducer)
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: gRPC service cannot create")
	}

	// Attach gRPC logger middleware for logging requests
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLoggerMiddleware)
	grpcServer := grpc.NewServer(grpcLogger)

	// Register the transactionService to the gRPC server
	pb.RegisterMoneyTransferServiceServer(grpcServer, tranService)

	// Enable reflection for gRPC, useful for debugging or using CLI tools like grpcurl
	reflection.Register(grpcServer)

	// Create a TCP listener for the gRPC server on the configured address
	listener, err := net.Listen("tcp", config.GRPCMoneyTransferServiceAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: gRPC service cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("MoneyTransfer Service: gRPC service started at %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("MoneyTransfer Service: gRPC service failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("MoneyTransfer Service: gRPC service shutdown gracefully")

		grpcServer.GracefulStop()
		log.Info().Msg("MoneyTransfer Service: gRPC server is stopped")

		return nil
	})
}

// runGatewayServer starts the gRPC Gateway server to serve HTTP requests, translating them into gRPC calls
func runGatewayServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	stores []db.Store,
	cache cache.Cache,
	taskProducer worker.TaskProducer,
) {
	// Create a new gRPC Gateway server
	tranService, err := gapi.NewService(config, stores, cache, taskProducer)
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: HTTPGateway service cannot create")
	}

	// Set custom JSON marshaling options for the gRPC Gateway
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true, // Use the proto field names in the JSON
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true, // Ignore unknown fields when unmarshaling
		},
	})

	// Create a new gRPC Gateway multiplexer to route HTTP requests
	grpcMux := runtime.NewServeMux(jsonOption)

	err = pb.RegisterMoneyTransferServiceHandlerServer(ctx, grpcMux, tranService)
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: HTTPGateway service cannot register handler")
	}

	// Create a new HTTP multiplexer for handling additional routes
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("MoneyTransfer Service: Cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/docs/", http.FileServer(statikFS))
	mux.Handle("/docs/", swaggerHandler)

	loggingHandler := logger.HttpLoggerMiddleware(mux)
	httpServer := &http.Server{
		Addr:    config.HTTPMoneyTransferServiceAddress,
		Handler: loggingHandler,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("MoneyTransfer Service: HTTPGateway service served at %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("MoneyTransfer Service: HTTPGateway service failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("MoneyTransfer Service: HTTPGateway service shutdown gracefully")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("MoneyTransfer Service: HTTPGateway service failed to shutdown ")
			return err
		}

		log.Info().Msg("MoneyTransfer Service: HTTPGateway service is stopped")
		return nil
	})
}
