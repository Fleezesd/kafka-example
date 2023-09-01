package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

var (
	addr          = flag.String("addr", ":8010", "The address to bind to")
	brokers       = flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
	verbose       = flag.Bool("verbose", false, "Turn on Sarama logging")
	certFile      = flag.String("certificate", "", "The optional certificate file for client authentication")
	keyFile       = flag.String("key", "", "The optional key file for client authentication")
	caFile        = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
	tlsSkipVerify = flag.Bool("tls-skip-verify", false, "Whether to skip TLS server cert verification")
)

func main() {
	flag.Parse()

	if *verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka broker: %s", strings.Join(brokerList, ","))

	// server init
	server := &Server{
		Datacollector:     newDataCollector(brokerList),
		AccessLogProducer: newAccessLogProducer(brokerList),
	}
	defer func() {
		if err := server.Close(); err != nil {
			log.Println("Failed to close server", err)
		}
	}()
	log.Fatal(server.Run(*addr))
}
