package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	cusdb "github.com/tunvx/simplebank/manage/db/sqlc"
	"github.com/tunvx/simplebank/notification/redis"
	"github.com/tunvx/simplebank/pkg/mail"
	"github.com/tunvx/simplebank/pkg/util"
	"golang.org/x/sync/errgroup"
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

	// If in development mode, configure logger to output to the console in a readable format
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// Create a new database connection pool
	connPool, err := pgxpool.New(context.Background(), config.DBSourceCoreDB)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer connPool.Close()

	// Create a new store to interact with the database
	cusStore := cusdb.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.DockerRedisAddress,
	}

	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, config, redisOpt, cusStore)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runTaskProcessor(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config,
	redisOpt asynq.RedisClientOpt,
	cusStore cusdb.Store,
) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := redis.NewRedisTaskProcessor(redisOpt, cusStore, mailer)

	log.Info().Msgf("start Task:Processor at [::]:%s", strings.Split(redisOpt.Addr, ":")[1])
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
