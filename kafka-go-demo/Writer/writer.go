package writer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func SimpleWriter() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "topic-A",
		// LeastBytes is a Balancer implementation that routes messages to the
		// partition that has received the least amount of data.
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close write:", err)
	}
}
