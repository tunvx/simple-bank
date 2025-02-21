package main

import (
	"io"
	"os"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tunvx/simplebank/common/util"
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
	runDBMigration(config.SourceSchemaURL, config.DBSourceOriginalDB)
}

// runDBMigration applies the database migrations to ensure the database schema is up-to-date
func runDBMigration(sourceSchemaURL string, dbSource string) {
	// Create a new migration instance
	migration, err := migrate.New(sourceSchemaURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	defer migration.Close()

	// Apply the migrations
	err = migration.Up()
	if err == migrate.ErrNoChange {
		log.Info().Msg("no migration changes detected")
	} else if err != nil {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	} else {
		log.Info().Msg("db migrated successfully")
	}
}