package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func processOrder(message *kafka.Message) error {
	order, err := orderFromMessage(message)
	if err != nil {
		return err
	}
	// Temporarily just log the order
	log.Printf("Build order %+v", order)
	return nil
}

func orderFromMessage(message *kafka.Message) (*Order, error) {
	order := &Order{}
	err := json.Unmarshal(message.Value, order)
	if err != nil {
		return nil, err
	}
	order.Status = "INITIAL"
	return order, nil
}

type processFunc func(*kafka.Message) error

func processAndLog(f processFunc, m *kafka.Message) {
	func(message *kafka.Message) {
		if err := f(message); err != nil {
			log.Printf("Encountered an error processing log: %v\n", err)
		}
	}(m)
}
