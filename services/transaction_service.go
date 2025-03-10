package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"payment-system/config"
	"payment-system/models"

	"github.com/segmentio/kafka-go"
)

// PublishTransaction sends a transaction to Kafka for processing
func PublishTransaction(transaction models.Transaction) error {
	// Convert transaction data to JSON
	data, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	message := kafka.Message{Value: data}

	// Publish to Kafka topic
	err = config.KafkaWriter.WriteMessages(context.Background(), message)
	if err != nil {
		log.Println("Kafka write error:", err)
		return errors.New("failed to queue transaction")
	}

	log.Println("Transaction published to Kafka successfully")
	return nil
}

// GetTransactionByID retrieves a transaction from Redis or MySQL
func GetTransactionByID(id uint) (*models.Transaction, error) {
	cacheKey := "transaction:" + string(rune(id))

	// Try fetching from Redis first
	cachedTxn, err := config.RDB.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var txn models.Transaction
		json.Unmarshal([]byte(cachedTxn), &txn)
		log.Println("Cache hit: Retrieved transaction from Redis")
		return &txn, nil
	}

	// If not found in cache, fetch from MySQL
	var transaction models.Transaction
	if err := config.DB.First(&transaction, id).Error; err != nil {
		return nil, errors.New("transaction not found")
	}

	// Store the transaction in Redis for future queries
	data, _ := json.Marshal(transaction)
	config.RDB.Set(config.Ctx, cacheKey, data, 10*time.Minute)

	log.Println("Cache miss: Retrieved transaction from MySQL and cached")
	return &transaction, nil
}
