package config

import (
	"log"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func InitKafka() {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("Kafka initialized")
}
