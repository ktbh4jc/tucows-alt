package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pgWriter, err := NewPostgresWriter()
	if err != nil {
		log.Fatal(err)
	}
	ws := NewWriterService(pgWriter)

	err = ws.Writer.Init()
	if err != nil {
		log.Fatal(err)
	}

	processor := NewProcessor(*ws)
	log.Printf("DB is up")

	// In full version I would build a consumer for each table
	orderConsumer, err := BuildConsumer("WRITE_ORDER")
	if err != nil {
		log.Fatal(err)
	}
	go RunConsumer(orderConsumer, *processor)
	// The below code prevents the service from closing until it receives a signal
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
