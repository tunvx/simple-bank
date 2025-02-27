package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const (
	TopicInternalTransferMoney = "topic_internal_transfer_money"
)

type PayloadInternalTransferMoney struct {
	SendingTranID   uuid.UUID 	`json:"sending_tran_id"`
	ReceivingTranID uuid.UUID 	`json:"receiving_tran_id"`
	Amount          int64  	  	`json:"amount"`
	CurrencyType    string 		`json:"currency_type"`
	SrcAccNumber    string 		`json:"src_acc_number"`
	SrcAccShardId 	int 		`json:"src_acc_shard_id"`
	BeneAccNumber   string 		`json:"bene_acc_number"`
	BeneAccShardId 	int 		`json:"bene_acc_shard_id"`
}

// SendInternalTransferMoneyMessageToKafka sends a message to Kafka with optional policies
func (kafkaProducer *KafkaTaskProducer) SendTaskInternalTransferMoney(
	ctx context.Context,
	payload *PayloadInternalTransferMoney,
) error {
	if kafkaProducer == nil {
        return fmt.Errorf("kafkaProducer is nil")
    }

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := &sarama.ProducerMessage{
		Topic: TopicInternalTransferMoney,
		Value: sarama.ByteEncoder(jsonPayload),
	}

	// Send message to Kafka
	kafkaProducer.kafkaProducer.Input() <- message
	log.Debug().Msgf("Kafka Producer: Sent InternalTransferMoney event with Payload: %+v", payload)


	select {
	case err := <-kafkaProducer.kafkaProducer.Errors():
		return fmt.Errorf("[Kafka] Failed to send internal transfer money message: %w", err)
	case msg := <-kafkaProducer.kafkaProducer.Successes():
		kafkaProducer.logger.Debug("[Kafka] Successfully sent MTT with SendingTranID ( %s ) to topic %s (Partition: %d, Offset: %d)", payload.SendingTranID, msg.Topic, msg.Partition, msg.Offset)
		return nil
	}
}

func (consumer *KafkaConsumer) ProcessTaskInternalTransferMoney(ctx context.Context, topic string) error {
    return nil
}