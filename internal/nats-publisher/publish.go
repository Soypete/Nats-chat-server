package natsbot

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

// Publish write message to nats server
func Publish(message string) {
	nc, err := nats.Connect("0.0.0.0:4222", nats.FlusherTimeout(time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nc.Publish(message, scanner.Bytes())
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
