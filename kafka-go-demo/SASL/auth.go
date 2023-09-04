package sasl

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Auth() {
	// plain SASL简单认证
	// mechanism := plain.Mechanism{
	// 	Username: "username",
	// 	Password: "password",
	// }

	ctx := context.Background()
	// scram 加密认证
	mechanism, err := scram.Mechanism(scram.SHA512, "username", "password")
	if err != nil {
		panic(err)
	}
	dialer := kafka.Dialer{
		Timeout: 10 * time.Second,
		// 当网络为“tcp”并且目标是同时具有 IPv4 和 IPv6 地址的主机名时，
		// DualStack 可以实现符合 RFC 6555 的“Happy Eyeballs”拨号。这允许客户端容忍一个地址族被悄悄破坏的网络
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	// connection
	conn, err := dialer.DialContext(ctx, "tcp", "localhost:9092")

	// reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "consumer-group-id",
		Topic:   "topic-A",
		Dialer:  &dialer,
	})

	// writer
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}
	w := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "topic-A",
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}

	// client
	client := kafka.Client{
		Addr:      kafka.TCP("localhost:9092"),
		Timeout:   10 * time.Second,
		Transport: sharedTransport,
	}
	fmt.Println(conn, r, w, client)
}
