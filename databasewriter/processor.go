package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Processor struct {
	WriterService WriterService
}

func NewProcessor(writerService WriterService) *Processor {
	return &Processor{
		WriterService: writerService,
	}
}

func (p *Processor) writeOrder(message *kafka.Message) (*int, error) {
	plainIntent, err := p.intentFromMessage(message)
	if err != nil {
		// If unable to extract intent, logs get a lot worse.
		// Something to consider if scaling.
		return nil, err
	}
	owc, err := p.owcFromMessage(message)
	if err != nil {
		return &plainIntent.IntentNumber, err
	}
	order := &Order{
		CustomerId: owc.CustomerId,
		ProductId:  owc.ProductId,
		Count:      owc.Count,
		Status:     owc.Status,
	}

	err = p.WriterService.Writer.CreateOrder(order)
	// no need to check for err as we return the same thing regardless
	return &plainIntent.IntentNumber, err
}

type processFunc func(*kafka.Message) (*int, error)

func (p *Processor) ProcessAndLog(f processFunc, m *kafka.Message) {
	func(message *kafka.Message) {
		if intent, err := f(message); err != nil {
			log.Printf("Encountered an error processing intent %d log: %v\n", &intent, err)
		} else {
			log.Printf("Processed intent %d", &intent)
		}
	}(m)
}

func (p *Processor) owcFromMessage(message *kafka.Message) (*OrderWithContext, error) {
	owc := &OrderWithContext{}
	err := json.Unmarshal(message.Value, owc)
	if err != nil {
		return nil, err
	}
	return owc, nil
}

func (p *Processor) intentFromMessage(message *kafka.Message) (*PlainIntent, error) {
	pi := &PlainIntent{}
	err := json.Unmarshal(message.Value, pi)
	if err != nil {
		return nil, err
	}
	return pi, nil
}
