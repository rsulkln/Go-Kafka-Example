package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	var ctx = context.Background()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
	})

	defer writer.Close()

	msg := []kafka.Message{
		{Key: []byte("user1"), Value: []byte(`{"order_id":1,"user_id":1}`)},
		{Key: []byte("user2"), Value: []byte(`{"order_id":2,"user_id":2}`)},
		{Key: []byte("user3"), Value: []byte(`{"order_id":3,"user_id":3}`)},
		{Key: []byte("user4"), Value: []byte(`{"order_id":4,"user_id":4}`)},
	}

	for _, m := range msg {
		if err := writer.WriteMessages(ctx, m); err != nil {
			log.Fatalln("you have an error to write msg kafka ", "error:", err)
		} else {
			log.Println("Message sent successfully!, message:", string(m.Value))
		}
	}
}
