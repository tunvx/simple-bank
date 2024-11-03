package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		options *OptionPolicy,
	) error
}

type KafkaTaskDistributor struct {
	client sarama.AsyncProducer // client to send tasks to Kafka
}

// NewKafkaTaskDistributor initializes a new Kafka task distributor
func NewKafkaTaskDistributor(brokers []string, kafkaConfig *sarama.Config) (TaskDistributor, error) {
	client, err := sarama.NewAsyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka async producer: %w", err)
	}

	return &KafkaTaskDistributor{
		client: client,
	}, nil
}

// Close gracefully shuts down the Kafka producer
func (distributor *KafkaTaskDistributor) Close() error {
	distributor.client.AsyncClose()
	return nil
}
