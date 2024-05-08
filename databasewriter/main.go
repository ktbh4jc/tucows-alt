package main

import "log"

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
}
