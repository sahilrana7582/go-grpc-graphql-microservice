package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafka(brokerURL, topic string) {
	writer = &kafka.Writer{
		Addr:         kafka.TCP(brokerURL),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	log.Println("Kafka producer initialized")
}

func PublishMessage(eventType string, payload interface{}) error {
	msgBytes, err := json.Marshal(map[string]interface{}{
		"type":    eventType,
		"payload": payload,
	})
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(eventType),
		Value: msgBytes,
		Time:  time.Now(),
	}

	return writer.WriteMessages(context.Background(), msg)

}
