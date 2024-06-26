package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "foo",
		"acks":              "all"})

	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	hp := NewHelloProducerImpl(p, "HELLO")
	op := NewOrderProducerImpl(p, "ORDER")
	server := NewAPIServer(":3000", hp, op)
	server.Run()
}
