package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	CustomerID int64 `json:"customer_id"`
	ShardID    int32 `json:"shard_id"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	// Create new task
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)

	// Enqueue task
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	// *** shard id
	shardId := payload.ShardID - 1

	customer, err := processor.stores[shardId].GetCustomerByID(ctx, payload.CustomerID)
	if err != nil {
		return fmt.Errorf("failed to get customer: %w", err)
	}

	randomUUID, _ := uuid.NewRandom()

	verifyEmail, err := processor.stores[shardId].CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		ID:           randomUUID,
		CustomerID:   customer.CustomerID,
		EmailAddress: customer.EmailAddress,
		SecretCode:   util.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	randomUUIDString, _ := util.ConvertUUIDToString(verifyEmail.ID)

	subject := "Welcome to Simple Bank"
	// TODO: replace this URL with an environment variable that points to a front-end page
	verifyUrl := fmt.Sprintf("http://localhost:82/v1/customers/verify_email/%s/%d/%s",
		randomUUIDString, payload.ShardID, verifyEmail.SecretCode)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, customer.FullName, verifyUrl)

	to := []string{customer.EmailAddress}
	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", customer.EmailAddress).Msg("processed task")
	return nil
}
