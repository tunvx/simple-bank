package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/management/db/sqlc"
)

const (
	TopicSendVerifyEmail = "topic:send_verify_email"
)

type PayloadSendVerifyEmail struct {
	CustomerRid string `json:"customer_rid"`
}

type OptionPolicy struct {
	Retry int
}

// DistributeTaskSendVerifyEmail sends a message to Kafka with optional policies
func (d *KafkaTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	options *OptionPolicy,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := &sarama.ProducerMessage{
		Topic: "topic:send_verify_email",
		Value: sarama.ByteEncoder(jsonPayload),
	}

	for i := 0; i <= options.Retry; i++ {
		// Send message to Kafka
		d.client.Input() <- message

		select {
		case err := <-d.client.Errors():
			if i == options.Retry {
				return fmt.Errorf("failed to send Kafka message after retries: %w", err)
			}
			fmt.Printf("Retrying Kafka message send: attempt %d/%d\n", i+1, options.Retry)
		case msg := <-d.client.Successes():
			log.Printf("Enqueued task to Kafka topic %s for customer %s", msg.Topic, payload.CustomerRid)
			log.Printf("Kafka message sent successfully to partition %d at offset %d", msg.Partition, msg.Offset)
			return nil
		}
	}
	return nil
}

func (processor *KafkaTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, topic string) error {
	partitions, err := processor.consumer.Partitions(topic)
	if err != nil {
		return fmt.Errorf("failed to get partitions for topic %s: %w", topic, err)
	}

	for _, partition := range partitions {
		pc, err := processor.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return fmt.Errorf("failed to start consumer for partition %d: %w", partition, err)
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()
			for msg := range pc.Messages() {
				if err := processor.handleMessage(ctx, msg); err != nil {
					log.Error().Err(err).Msg("error processing message")
				}
			}
		}(pc)
	}

	// Wait for the context to be done before shutting down
	<-ctx.Done()

	return nil
}

// handleMessage processes a single Kafka message
func (processor *KafkaTaskProcessor) handleMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal message payload: %w", err)
	}

	// Retrieve customer information
	customer, err := processor.cusStore.GetCustomerByRid(ctx, payload.CustomerRid)
	if err != nil {
		return fmt.Errorf("failed to get customer: %w", err)
	}

	// Create and send verification email
	verifyEmail, err := processor.cusStore.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		SecretCode: util.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	subject := "Welcome to Simple Bank"
	verifyUrl := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>`,
		customer.Fullname, verifyUrl)

	to := []string{customer.Email}
	if err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	log.Info().Str("email", customer.Email).Msg("processed email verification task")
	return nil
}
