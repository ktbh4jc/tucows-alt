package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderProcessor struct {
	producer     *kafka.Producer
	topic        string
	deliveryChan chan kafka.Event
}

func NewOrderProcessor(producer *kafka.Producer, topic string) *OrderProcessor {
	return &OrderProcessor{
		producer:     producer,
		topic:        topic,
		deliveryChan: make(chan kafka.Event),
	}
}

func (op OrderProcessor) ProcessOrder(message *kafka.Message) error {
	order, err := orderFromMessage(message)
	if err != nil {
		return err
	}
	// Temporarily just log the order
	log.Printf("Build order %+v", order)
	err = op.PublishOrder(*order)
	return err
}

func orderFromMessage(message *kafka.Message) (*Order, error) {
	order := &Order{}
	err := json.Unmarshal(message.Value, order)
	if err != nil {
		return nil, err
	}
	order.Status = "IN PROGRESS"
	return order, nil
}

type processFunc func(*kafka.Message) error

func (op OrderProcessor) ProcessAndLog(f processFunc, m *kafka.Message) {
	func(message *kafka.Message) {
		if err := f(message); err != nil {
			log.Printf("Encountered an error processing log: %v\n", err)
		}
	}(m)
}

func (op OrderProcessor) PublishOrder(order Order) error {
	payload, err := json.Marshal(order)
	if err != nil {
		return err
	}
	message := buildMessage("WRITE_ORDER", payload)
	err = op.producer.Produce(message, op.deliveryChan)
	if err != nil {
		return err
	}
	log.Printf("Sending order %s\n", string(payload))
	<-op.deliveryChan
	return nil
}

func buildMessage(topic string, payload []byte) *kafka.Message {
	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload}
}
