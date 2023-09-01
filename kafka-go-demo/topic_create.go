package main

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func TopicCreate() {
	// to create topic when auto.create.topics.enable='false'
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	// controller 向 kafka 请求当前controller并返回其 URL信息 即broker
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	// topic config
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1, // 分区个数
			ReplicationFactor: 1, // 分区副本
		},
	}
	// 创建topic		leader create topic
	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
