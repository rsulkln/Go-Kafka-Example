package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "orders",
		GroupID:     "order-consumer-group-test3",
		StartOffset: kafka.LastOffset,
	})
	defer reader.Close()

	log.Println("reader two started...!")

	for {
		var ctx, cancle = context.WithTimeout(context.Background(), 5*time.Second)
		msg, err := reader.ReadMessage(ctx)
		cancle()

		if err != nil {
			log.Println("No new message...")
			time.Sleep(1 * time.Second)
			continue
		}
		log.Printf("Message received: partition=%d offset=%d key=%s value=%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}

}
