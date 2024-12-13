package main

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natefinch/lumberjack"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pb "github.com/tunvx/simplebank/grpc/pb/transactions"
	db "github.com/tunvx/simplebank/manage/db/sqlc"
	worker "github.com/tunvx/simplebank/notification/redis"
	"github.com/tunvx/simplebank/pkg/logger"
	"github.com/tunvx/simplebank/pkg/util"
	"github.com/tunvx/simplebank/transactions/cache"
	"github.com/tunvx/simplebank/transactions/gapi"
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

	// Set up lumberjack for log rotation
	logFile := &lumberjack.Logger{
		Filename:   "/var/log/service.log",
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

	// Database connection pool configuration
	connPoolConfig, err := pgxpool.ParseConfig(config.DBSourceCoreDB)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse database config")
	}
	// connPoolConfig.MaxConns = 150
	// connPoolConfig.MinConns = 50
	// connPoolConfig.MaxConnLifetime = time.Minute * 10
	// connPoolConfig.MaxConnIdleTime = time.Minute * 2

	// Create PostgreSQL connection pool
	connPool, err := pgxpool.NewWithConfig(context.Background(), connPoolConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer connPool.Close()

	// Create a new store to interact with the database
	store := db.NewStore(connPool)

	// Redis connection pool configuration
	redisOpt1 := redis.Options{
		Addr:            config.InternalRedisAddress,
		MaxActiveConns:  120,
		ConnMaxIdleTime: time.Minute * 5,
	}
	cache := cache.NewRedisCache(&redisOpt1)

	redisOpt2 := asynq.RedisClientOpt{
		Addr: config.InternalRedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt2)
	log.Info().Msgf("start Task:Distributor at :: %s", redisOpt2.Addr)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runGrpcServer(ctx, waitGroup, config, store, cache, taskDistributor)
	runGatewayServer(ctx, waitGroup, config, store, cache, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}

	// log.Info().Msg("add trigger to run github action")
}

// runGrpcServer starts the gRPC server for handling core banking services
func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
	cache cache.Cache,
	taskDistributor worker.TaskDistributor,
) {
	// Create a new transaction service
	tranService, err := gapi.NewService(config, store, cache, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create transaction service")
	}

	// Attach gRPC logger middleware for logging requests
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	// Register the transactionService to the gRPC server
	pb.RegisterTransactionServiceServer(grpcServer, tranService)

	// Enable reflection for gRPC, useful for debugging or using CLI tools like grpcurl
	reflection.Register(grpcServer)

	// Create a TCP listener for the gRPC server on the configured address
	listener, err := net.Listen("tcp", config.GRPCTransactionServiceAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC Transaction:Service as server at :: %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gRPC server")

		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server is stopped")

		return nil
	})
}

// runGatewayServer starts the gRPC Gateway server to serve HTTP requests, translating them into gRPC calls
func runGatewayServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
	cache cache.Cache,
	taskDistributor worker.TaskDistributor,
) {
	// Create a new gRPC Gateway server
	tranService, err := gapi.NewService(config, store, cache, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
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

	err = pb.RegisterTransactionServiceHandlerServer(ctx, grpcMux, tranService)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	// Create a new HTTP multiplexer for handling additional routes
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Serve Swagger documentation for the API at /swagger/
	fs := http.FileServer(http.Dir("./doc/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	// Register pprof handlers for profiling
	// mux.Handle("/debug/pprof/", http.DefaultServeMux)

	loggingHandler := logger.HttpLogger(mux)
	httpServer := &http.Server{
		Addr:    config.HTTPTransactionServiceAddress,
		Handler: loggingHandler,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start HTTPGateway Transaction:Service as server at :: %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("HTTP gateway server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown HTTP gateway server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown HTTP gateway server")
			return err
		}

		log.Info().Msg("HTTP gateway server is stopped")
		return nil
	})
}
