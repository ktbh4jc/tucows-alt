package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Todo: pull this from .env
	topic := "HELLO"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "bar",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		event := consumer.Poll(100)
		switch e := event.(type) {
		case *kafka.Message:
			log.Printf("processing hello: %s\n", string(e.Value))
		case *kafka.Error:
			log.Printf("%v\n", e)
		}
	}
}
