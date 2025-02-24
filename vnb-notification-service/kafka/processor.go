package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/tunvx/simplebank/common/mail"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
)

type TaskProcessor interface {
	Start(ctx context.Context) error
	Shutdown() error
	ProcessTaskSendVerifyEmail(ctx context.Context, payload []byte) error
}

type KafkaTaskProcessor struct {
	consumer sarama.Consumer
	cusStore db.Store
	mailer   mail.EmailSender
}

func NewKafkaTaskProcessor(brokers []string, groupID string, cusStore db.Store, mailer mail.EmailSender) (*KafkaTaskProcessor, error) {
	// consumer, err := sarama.NewConsumerFromClient(brokers, nil)
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	return &KafkaTaskProcessor{
		consumer: consumer,
		cusStore: cusStore,
		mailer:   mailer,
	}, nil
}

// Close shuts down the Kafka consumer gracefully
func (processor *KafkaTaskProcessor) Close() error {
	return processor.consumer.Close()
}
