package worker

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
)

type TaskConsumer interface {
	Start(ctx context.Context) error
	Shutdown() error
	ProcessTaskInternalTransferMoney(ctx context.Context, payload []byte) error
}

type KafkaConsumer struct {
	kafkaConsumer sarama.Consumer
	stores        []db.Store
}

func NewKafkaConsumer(brokers []string, groupID string, stores []db.Store) (*KafkaConsumer, error) {
	logger := NewLogger()
	sarama.Logger = logger
	
	// consumer, err := sarama.NewConsumerFromClient(brokers, nil)
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating Kafka consumer: %w", err)
	}

	return &KafkaConsumer{
		kafkaConsumer: consumer,
		stores: stores,
	}, nil
}

// Close shuts down the Kafka consumer gracefully
func (comsumer *KafkaConsumer) Close() error {
	return comsumer.kafkaConsumer.Close()
}
