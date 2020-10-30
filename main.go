package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("Test-message", func(m *nats.Msg) {
		// wg.Done()
		fmt.Println(string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}
	// Do something with the connection
	err = nc.Publish("Test-message", []byte("hello server"))
	if err != nil {
		log.Fatal(err)
	}
	// Wait for a message to come in
	wg.Wait()
}
