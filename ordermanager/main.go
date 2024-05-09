package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Todo: pull this from .env
	topic := "ORDER"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "bar",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatal(err)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "foo",
		"acks":              "all"})

	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	op := NewOrderProcessor(p, "WRITE_ORDER")

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		event := consumer.Poll(100)
		switch e := event.(type) {
		case *kafka.Message:
			log.Printf("processing order: %s\n", string(e.Value))
			op.ProcessAndLog(op.ProcessOrder, e)
		case *kafka.Error:
			log.Printf("%v\n", e)
		}
	}
}
