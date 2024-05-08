package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type HelloProducer interface {
	SayHello() error
}

type HelloProducerImpl struct {
	producer     *kafka.Producer
	topic        string
	deliveryChan chan kafka.Event
}

func NewHelloProducerImpl(p *kafka.Producer, topic string) *HelloProducerImpl {
	return &HelloProducerImpl{
		producer:     p,
		topic:        topic,
		deliveryChan: make(chan kafka.Event, 10000),
	}
}

func (hp *HelloProducerImpl) SayHello() error {
	payload := []byte("Hello world")
	message := buildMessage(hp.topic, payload)

	err := hp.producer.Produce(message, hp.deliveryChan)
	if err != nil {
		return err
	}

	log.Printf("Sending payload %s\n", string(payload))
	<-hp.deliveryChan
	return nil
}

type OrderProducer interface {
	PlaceOrder(OrderRequest) error
}

type OrderProducerImpl struct {
	producer     *kafka.Producer
	topic        string
	deliveryChan chan kafka.Event
}

func NewOrderProducerImpl(p *kafka.Producer, topic string) *OrderProducerImpl {
	return &OrderProducerImpl{
		producer:     p,
		topic:        topic,
		deliveryChan: make(chan kafka.Event, 10000),
	}
}

func (op *OrderProducerImpl) PlaceOrder(request OrderRequest) error {
	payload, err := json.Marshal(request)
	if err != nil {
		return err
	}
	message := buildMessage(op.topic, payload)
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
