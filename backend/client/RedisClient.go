package client

import (
	"accounting/types"
	"encoding/json"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		client: rdb,
	}
}

func (c *RedisClient) AddTransaction(transaction types.TransactionInput) error {
	payloadBytes, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Store transaction in a Redis list
	if err := c.client.RPush(ctx, "transactions", payloadBytes).Err(); err != nil {
		return fmt.Errorf("error storing transaction in Redis list: %v", err)
	}

	return nil
}

func (c *RedisClient) GetTransactions(page, perPage int) ([]types.Transaction, error) {
	// Calculate the range for pagination
	start := (page - 1) * perPage
	end := start + perPage - 1

	// Fetch transactions from the Redis list
	data, err := c.client.LRange(ctx, "transactions", int64(start), int64(end)).Result()
	if err != nil {
		return nil, fmt.Errorf("error fetching transactions from Redis list: %v", err)
	}

	var transactions []types.Transaction
	for _, item := range data {
		var transaction types.Transaction
		if err := json.Unmarshal([]byte(item), &transaction); err != nil {
			return nil, fmt.Errorf("error unmarshalling transaction data: %v", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (c *RedisClient) GetTransactionCount() int {
	// Get the length of the Redis list
	length, err := c.client.LLen(ctx, "transactions").Result()
	if err != nil {
		log.Println("error getting list length from Redis:", err)
		return 0
	}
	return int(length)
}
