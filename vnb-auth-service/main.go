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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natefinch/lumberjack"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	db "github.com/tunvx/simplebank/authsrv/db/sqlc"
	"github.com/tunvx/simplebank/authsrv/gapi"
	"github.com/tunvx/simplebank/common/logger"
	"github.com/tunvx/simplebank/common/util"
	_ "github.com/tunvx/simplebank/grpc/doc/auth/statik"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
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
		log.Fatal().Err(err).Msg("Auth Service: cannot load config")
	}

	// Set up lumberjack for log rotation
	logFile := &lumberjack.Logger{
		Filename:   "/var/log/service-auth.log",
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
		log.Fatal().Msgf("Auth Service: Unknown environment ( %s ) ", config.Environment)
	}

	// Run database migrations
	runDBMigration(config.SourceSchemaURL, config.DBSourceAuthDB)

	// Create a new database connection pool
	connPool, err := pgxpool.New(context.Background(), config.DBSourceAuthDB)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Cannot established connection pool to database")
	}
	defer connPool.Close()

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// Create a new store to interact with the database
	store := db.NewStore(connPool)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runGatewayServer(ctx, waitGroup, config, store)
	runGrpcServer(ctx, waitGroup, config, store)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Error from wait group")
	}

	// log.Info().Msg("add trigger to run github action")
}

// runDBMigration applies the database migrations to ensure the database schema is up-to-date
func runDBMigration(sourceSchemaURL string, dbSource string) {
	// Create a new migration instance
	migration, err := migrate.New(sourceSchemaURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Cannot create new migrate instance")
	}

	defer migration.Close()

	// Apply the migrations
	err = migration.Up()
	if err == migrate.ErrNoChange {
		log.Info().Msg("Auth Service: No migration changes detected")
	} else if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Failed to run migrate up")
	} else {
		log.Info().Msg("Auth Service: Database migrated successfully")
	}
}

// runGrpcServer starts the gRPC server for handling core banking services
func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,

) {
	// Create a new Auth service
	authService, err := gapi.NewService(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: cannot create new gRPC instance of service")
	}

	// Attach gRPC logger, icall middleware for requests
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLoggerMiddleware)
	grpcServer := grpc.NewServer(grpcLogger)

	// Register the AuthService to the gRPC server
	pb.RegisterAuthServer(grpcServer, authService)

	// Enable reflection for gRPC, useful for debugging or using CLI tools like grpcurl
	reflection.Register(grpcServer)

	// Create a TCP listener for the gRPC server on the configured address
	listener, err := net.Listen("tcp", config.GRPCAuthServiceAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("Auth Service: gRPC Service started on %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("Auth Service: gRPC Service failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Auth Service: Graceful shutdown gRPC server")

		grpcServer.GracefulStop()
		log.Info().Msg("Auth Service: gRPC service is stopped")

		return nil
	})
}

// runGatewayServer starts the gRPC Gateway server to serve HTTP requests,
// translating them into gRPC calls
func runGatewayServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
) {
	// Create a new gRPC Gateway server
	authService, err := gapi.NewService(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: cannot create new HTTP instance of service")
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

	err = pb.RegisterAuthHandlerServer(ctx, grpcMux, authService)
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Cannot register handler of service")
	}

	// Create a new HTTP multiplexer for handling additional routes
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Auth Service: Cannot create statik fs for swagger")
	}

	swaggerHandler := http.StripPrefix("/docs/", http.FileServer(statikFS))
	mux.Handle("/docs/", swaggerHandler)

	loggingHandler := logger.HttpLoggerMiddleware(mux)
	httpServer := &http.Server{
		Handler: loggingHandler,
		Addr:    config.HTTPAuthServiceAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("Auth Service: HTTP Service started on %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("Auth Service: HTTP gateway service failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Auth Service: Graceful shutdown HTTP gateway server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Auth Service: Failed to shutdown HTTP gateway server")
			return err
		}

		log.Info().Msg("Auth Service: HTTP gateway server is stopped")
		return nil
	})
}
