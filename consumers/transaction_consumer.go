package consumers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"payment-system/config"
	"payment-system/models"

	"github.com/segmentio/kafka-go"
)

// StartTransactionConsumer listens for transaction events from Kafka and processes them
func StartTransactionConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
		GroupID: "transaction_group",
	})
	defer reader.Close()

	log.Println("Transaction consumer started...")

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading Kafka message:", err)
			continue
		}

		var transaction models.Transaction
		if err := json.Unmarshal(message.Value, &transaction); err != nil {
			log.Println("Error decoding Kafka message:", err)
			continue
		}

		log.Printf("Processing transaction: %+v\n", transaction)

		// Simulate transaction processing time
		time.Sleep(2 * time.Second)

		// Update the transaction as completed
		transaction.Status = "Completed"
		config.DB.Save(&transaction)

		// Store processed transaction in Redis for fast retrieval
		data, _ := json.Marshal(transaction)
		cacheKey := "transaction:" + string(rune(transaction.ID))
		config.RDB.Set(config.Ctx, cacheKey, data, 10*time.Minute)

		log.Println("Transaction processed successfully and stored in Redis")
	}
}
