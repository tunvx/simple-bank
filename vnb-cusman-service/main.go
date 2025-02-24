package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tunvx/simplebank/common/logger"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	"github.com/tunvx/simplebank/cusmansrv/gapi"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman"
	"github.com/tunvx/simplebank/notificationsrv/redis"
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
		Filename:   "/var/log/service-management.log",
		MaxSize:    1,    // Maximum size in MB before rotation
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
		log.Fatal().Msgf("Unknown environment: %v", config.Environment)
	}

	// Run database migrations
	runDBMigration(config.SourceSchemaURL, config.ListDBSourceCoreDB, config.NumCoreDBShard)

	// Establish SQL Store for all shards
	stores, err := establishShardedSQLStore(config.ListDBSourceCoreDB, config.NumCoreDBShard)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize sharded stores")
	}
	defer func() {
		for _, store := range stores {
			if sqlStore, ok := store.(*db.SQLStore); ok {
				sqlStore.Close()
			}
		}
	}()

	redisOpt := asynq.RedisClientOpt{
		Addr: config.InternalRedisAddress,
	}

	taskDistributor := redis.NewRedisTaskDistributor(redisOpt)
	log.Info().Msgf("start Task:Distributor at :: %s", redisOpt.Addr)

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	runGrpcServer(ctx, waitGroup, config, stores, taskDistributor)
	runGatewayServer(ctx, waitGroup, config, stores, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}

	// log.Info().Msg("add trigger to run github action")
}

func establishShardedSQLStore(listShardDatabaseURL []string, numShard int) ([]db.Store, error) {
	// Check if the length of listShardDatabaseURL matches numShard
	if len(listShardDatabaseURL) != numShard {
		log.Fatal().Msgf("The length of listShardDatabaseURL [ %d ] does not match numShard [ %d ]", len(listShardDatabaseURL), numShard)
	}

	var stores []db.Store

	// Loop over each shard and apply migration
	for shardID := 1; shardID <= numShard; shardID++ {
		// Access the shard database URL based on shardID
		shardDatabaseURL := listShardDatabaseURL[shardID-1]
		// *************************************************************
		connPoolConfig, err := pgxpool.ParseConfig(shardDatabaseURL)
		if err != nil {
			log.Error().Err(err).Msgf("unable to parse database config, error: %v", err)
			continue
		}

		connPoolConfig.MinConns = 40
		connPoolConfig.MaxConns = 100
		connPoolConfig.MaxConnLifetime = time.Minute * 10
		connPoolConfig.MaxConnIdleTime = time.Minute * 2

		// Create pool with custom configuration
		connPool, err := pgxpool.NewWithConfig(context.Background(), connPoolConfig)
		if err != nil {
			log.Error().Err(err).Msgf("cannot connect to shard [ %d ], error: %v", err)
			continue
		} else {
			log.Info().Msgf("successfully created pool connection to shard [ %d ]", shardID)
		}

		// Create a new store to interact with the database
		store := db.NewStore(connPool)
		stores = append(stores, store)
	}

	if len(stores) == 0 {
		return nil, fmt.Errorf("failed to initialize any database shard")
	}

	return stores, nil
}

// runDBMigration applies the database migrations to ensure the database schema is up-to-date
func runDBMigration(sourceSchemaURL string, listShardDatabaseURL []string, numShard int) {
	// Check if the length of listShardDatabaseURL matches numShard
	if len(listShardDatabaseURL) != numShard {
		log.Fatal().Msgf("The length of listShardDatabaseURL [ %d ] does not match numShard [ %d ]", len(listShardDatabaseURL), numShard)
	}

	// Loop over each shard and apply migration
	for shardID := 1; shardID <= numShard; shardID++ {
		// Access the shard database URL based on shardID
		shardDatabaseURL := listShardDatabaseURL[shardID-1]

		conn, err := sql.Open("postgres", shardDatabaseURL)
		if err != nil {
			log.Error().Err(err).Msgf("cannot connect to shard [ %d ], error: %v", shardID, err)
			continue // Skip the current shard and continue with the next one
		} else {
			log.Info().Msgf("Successfully connected to shard [ %d ]", shardID)
		}
		defer conn.Close()

		// Create id schema, sequence, and function for id generator before migration
		const idGeneratorSchema = "shard_id_generator"
		createInitSchemaSQL := fmt.Sprintf(`
			CREATE SCHEMA IF NOT EXISTS %s;
			CREATE SEQUENCE IF NOT EXISTS %s.shard_id_sequence;

			CREATE OR REPLACE FUNCTION %s.generate_id(OUT result bigint) AS $$
			DECLARE
				our_epoch bigint := 1314220021721;
				seq_id bigint;
				now_millis bigint;
			BEGIN
				SELECT nextval('%s.shard_id_sequence') %% 4096 INTO seq_id;
				SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
				
				-- sign 1 bit, timestamp 42 bits, shard_id 9 bits, seq_id 12 bits
				result := (now_millis - our_epoch) << 21;	-- Shift left by 21 bits to make room for shard_id + sequence_id
				result := result | (%d << 12);				-- Shift left by 12 bits to make room for sequence_id
				result := result | (seq_id); 				-- Keep sequence_id in the last 12 bits                      
			END;
			$$ LANGUAGE PLPGSQL;
		`, idGeneratorSchema, idGeneratorSchema, idGeneratorSchema, idGeneratorSchema, shardID)

		_, err = conn.Exec(createInitSchemaSQL)
		if err != nil {
			log.Error().Err(err).Msgf("failed to create ID generator function for shard [ %d ], error: %v", shardID, err)
			continue // Skip the current shard and continue with the next one
		}
		log.Info().Msgf("ID generator function created successfully for shard [ %d ]", shardID)

		// Init migration
		migration, err := migrate.New(sourceSchemaURL, shardDatabaseURL)
		if err != nil {
			log.Error().Err(err).Msgf("cannot create migrate instance for shard [ %d ], error: %v", shardID, err)
			continue // Skip the current shard and continue with the next one
		}
		defer migration.Close()

		// Apply (run) the migration
		err = migration.Up()
		if err == migrate.ErrNoChange {
			log.Info().Msgf("no migration changes detected for shard [ %d ]", shardID)
		} else if err != nil {
			log.Error().Err(err).Msgf("failed to run migrate up for shard [ %d ], error: %v", shardID, err)
			continue // Skip the current shard and continue with the next one
		} else {
			log.Info().Msgf("db migrated successfully for shard [ %d ]", shardID)
		}
	}
}

// runGrpcServer starts the gRPC server for handling core banking services
func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	stores []db.Store,
	taskDistributor redis.TaskDistributor,
) {
	// Create a new manage service
	manageService, err := gapi.NewService(config, stores, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create manage service")
	}

	// Attach gRPC logger middleware for logging requests
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	// Register the manageService to the gRPC server
	pb.RegisterCustomerManagementServiceServer(grpcServer, manageService)

	// Enable reflection for gRPC, useful for debugging or using CLI tools like grpcurl
	reflection.Register(grpcServer)

	// Create a TCP listener for the gRPC server on the configured address
	listener, err := net.Listen("tcp", config.GRPCManageServiceAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC Manage:Service as server at :: %s", listener.Addr().String())

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
	stores []db.Store,
	taskDistributor redis.TaskDistributor,
) {
	// Create a new gRPC Gateway server
	manageService, err := gapi.NewService(config, stores, taskDistributor)
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

	err = pb.RegisterCustomerManagementServiceHandlerServer(ctx, grpcMux, manageService)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	// Create a new HTTP multiplexer for handling additional routes
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Serve Swagger documentation for the API at /swagger/
	fs := http.FileServer(http.Dir("./doc/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	loggingHandler := logger.HttpLogger(mux)
	httpServer := &http.Server{
		Handler: loggingHandler,
		Addr:    config.HTTPManageServiceAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start HTTPGateway Manage:Service as server at :: %s", httpServer.Addr)
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
