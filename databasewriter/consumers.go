package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func BuildConsumer(topic string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "baz",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func RunConsumer(consumer *kafka.Consumer, processor Processor) {
	for {
		event := consumer.Poll(100)
		switch e := event.(type) {
		case *kafka.Message:
			log.Printf("processing message: %s\n", string(e.Value))
			processor.ProcessAndLog(processor.writeOrder, e)
		case *kafka.Error:
			log.Printf("%v\n", e)
		}
	}
}
