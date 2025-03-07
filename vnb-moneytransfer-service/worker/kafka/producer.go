package worker

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

type TaskProducer interface {
	SendTaskInternalTransferMoney(
		ctx context.Context,
		payload *PayloadInternalTransferMoney,
	) error
}

type KafkaTaskProducer struct {
	kafkaProducer sarama.AsyncProducer // client to send tasks to Kafka
	logger 		  Logger
}

// NewKafkaTaskProducer initializes a new Kafka task distributor
func NewKafkaTaskProducer(brokers []string, topics []string, kafkaConfig *sarama.Config) (TaskProducer, error) {
	logger := NewLogger()
	sarama.Logger = logger

	// Create Kafka Admin client
	admin, err := sarama.NewClusterAdmin(brokers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka admin client: %w", err)
	}
	defer admin.Close() 
	
	// Create topics if they don't exist
	for _, topic := range topics {
		topicDetail := &sarama.TopicDetail{
			NumPartitions:     8,
			ReplicationFactor: 1,
			ConfigEntries:     make(map[string]*string),
		}

		err := admin.CreateTopic(topic, topicDetail, false)
		if err != nil {
			if topicErr, ok := err.(*sarama.TopicError); ok && topicErr.Err == sarama.ErrTopicAlreadyExists {
				log.Printf("kafka producer: topic ( %s ) already exists", topic)
				continue
			}
			return nil, fmt.Errorf("kafka producer: failed to create topic ( %s ): %w", topic, err)
		}
	}

	// Create Kafka producer
	newKafkaProducer, err := sarama.NewAsyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka async producer: %w", err)
	}

	return &KafkaTaskProducer{
		kafkaProducer: newKafkaProducer,
	}, nil
}

// Close gracefully shuts down the Kafka producer
func (producer *KafkaTaskProducer) Close() error {
	if producer.kafkaProducer != nil {
		producer.kafkaProducer.AsyncClose()
	}
	return nil
}
