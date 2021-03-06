package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"b-2.tlio-prd-sg-msk-c.s18psb.c2.kafka.ap-southeast-1.amazonaws.com:9092,b-1.tlio-prd-sg-msk-c.s18psb.c2.kafka.ap-southeast-1.amazonaws.com:9092"},
		Topic:       "audit-prd",
		Partition:   0,
		GroupID:     "TestGroup",
		MaxAttempts: 3,
		MaxWait:     1 * time.Second,
		MinBytes:    1e1,  // 10KB
		MaxBytes:    10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("%v", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
