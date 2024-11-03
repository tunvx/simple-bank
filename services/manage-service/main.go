package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pb "github.com/tunvx/simplebank/grpc/pb/manage"
	db "github.com/tunvx/simplebank/manage/db/sqlc"
	"github.com/tunvx/simplebank/manage/gapi"
	"github.com/tunvx/simplebank/pkg/logger"
	"github.com/tunvx/simplebank/pkg/mail"
	"github.com/tunvx/simplebank/pkg/util"
	"github.com/tunvx/simplebank/worker"
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
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	// If in development mode, configure logger to output to the console in a readable format
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPoolConfig, err := pgxpool.ParseConfig(config.DBSourceCoreDB)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse database config")
	}
	connPoolConfig.MaxConns = 100
	connPoolConfig.MinConns = 40
	connPoolConfig.MaxConnLifetime = time.Minute * 10
	connPoolConfig.MaxConnIdleTime = time.Minute * 2

	// Tạo pool với cấu hình đã điều chỉnh
	connPool, err := pgxpool.NewWithConfig(context.Background(), connPoolConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer connPool.Close()

	// Run database migrations
	runDBMigration(config.MigrationURL, config.DBSourceCoreDB)

	// Create a new store to interact with the database
	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, config, redisOpt, store)
	runGrpcServer(ctx, waitGroup, config, store, taskDistributor)
	runGatewayServer(ctx, waitGroup, config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

// runDBMigration applies the database migrations to ensure the database schema is up-to-date
func runDBMigration(migrationURL string, dbSource string) {
	// Create a new migration instance
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	// Apply the migrations (skip if no changes are detected)
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
	defer migration.Close()
}

func runTaskProcessor(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	redisOpt asynq.RedisClientOpt,
	store db.Store,
) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)

	log.Info().Msgf("start task processor at [::]:%s", strings.Split(redisOpt.Addr, ":")[1])
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown task processor")

		taskProcessor.Shutdown()
		log.Info().Msg("task processor is stopped")

		return nil
	})
}

// runGrpcServer starts the gRPC server for handling core banking services
func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
	taskDistributor worker.TaskDistributor,
) {
	// Create a new account service
	accountService, err := gapi.NewService(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create account service")
	}

	// Attach gRPC logger middleware for logging requests
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	// Register the accountService to the gRPC server
	pb.RegisterManageServiceServer(grpcServer, accountService)

	// Enable reflection for gRPC, useful for debugging or using CLI tools like grpcurl
	reflection.Register(grpcServer)

	// Create a TCP listener for the gRPC server on the configured address
	listener, err := net.Listen("tcp", config.GRPCManageServiceAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC ManageService as server at %s", listener.Addr().String())

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
	taskDistributor worker.TaskDistributor,
) {
	// Create a new gRPC Gateway server
	accountService, err := gapi.NewService(config, store, taskDistributor)
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

	err = pb.RegisterManageServiceHandlerServer(ctx, grpcMux, accountService)
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
		log.Info().Msgf("start HTTPGateway ManageService as server at [::]:%s", strings.Split(httpServer.Addr, ":")[1])
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
