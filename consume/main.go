package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	// Subscribe
	sub, err := nc.SubscribeSync("Test-message")
	if err != nil {
		log.Fatal(err)
	}
	// Wait for a message
	msg, err := sub.NextMsg(10 * time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)
}
