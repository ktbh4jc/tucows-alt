package main

import (
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
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &hp.topic, Partition: kafka.PartitionAny},
		Value:          payload}

	err := hp.producer.Produce(message, hp.deliveryChan)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sending payload %v\n", payload)
	<-hp.deliveryChan
	return nil
}
