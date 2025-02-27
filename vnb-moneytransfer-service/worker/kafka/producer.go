package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
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
	logger := log.Default()
	sarama.Logger = logger

	// Create Kafka Admin client
	admin, err := sarama.NewClusterAdmin(brokers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka admin client: %w", err)
	}
	defer admin.Close() // Đóng admin khi kết thúc hàm
	
	// Create topics if they don't exist
	for _, topic := range topics {
		topicDetail := &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
			ConfigEntries:     make(map[string]*string),
		}

		err := admin.CreateTopic(topic, topicDetail, false)
		if err != nil {
			// Kiểm tra chính xác lỗi topic đã tồn tại
			if topicErr, ok := err.(*sarama.TopicError); ok && topicErr.Err == sarama.ErrTopicAlreadyExists {
				logger.Printf("Topic ( %s ) already exists", topic)
				continue
			}
			return nil, fmt.Errorf("failed to create topic ( %s ): %w", topic, err)
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
