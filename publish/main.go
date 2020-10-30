package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with the connection
	err = nc.Publish("Test-message", []byte("hello server"))
	if err != nil {
		log.Fatal(err)
	}
}
