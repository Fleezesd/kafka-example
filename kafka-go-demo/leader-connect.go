package main

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// to connect to the kafka leader via an existing non-leader connection
func LeaderConnect() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller() // broker
	if err != nil {
		panic(err.Error())
	}
	var connLeader *kafka.Conn
	connLeader, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}

	defer connLeader.Close()

	//跟topic_create 类似先连接到leader connection
}
