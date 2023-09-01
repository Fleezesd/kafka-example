package reader

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// 单topic partition
func SimpleReader() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   "consumer-group-id",
		Topic:     "my-topic",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	// offset deadline 与reader的offset不同
	// r.SetOffset(10)

	ctx := context.Background()
	for {
		// m, err := r.ReadMessage(context.Background())
		m, err := r.FetchMessage(ctx) // 显示提交
		fmt.Println(m)
		if err != nil {
			break
		}
		fmt.Printf("meesage at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messaages:", err)
		}
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
