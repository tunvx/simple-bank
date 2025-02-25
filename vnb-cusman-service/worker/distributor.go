package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client // use client to send tasks to redis queue
}

func NewRedisTaskDistributor(redisOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpts)
	return &RedisTaskDistributor{
		client: client,
	}
}
