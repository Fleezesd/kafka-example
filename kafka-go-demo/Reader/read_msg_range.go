package reader

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func logf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
}

func RangeReadMsg() {
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	batchSize := int(10e6) // 10MB

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:       "my-topic1",
		Partition:   0,
		MaxBytes:    batchSize,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	})
	// 设置起始读取信息的时间 offsetAt
	r.SetOffsetAt(context.Background(), startTime)

	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			break
		}
		if m.Time.After(endTime) {
			break
		}
		// TODO: process message
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
