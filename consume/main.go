package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("0.0.0.0:4222")
	nc.Flush()
	defer nc.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nc.ConnectedServerId())
	// Subscribe
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("Test-message", func(m *nats.Msg) {
		// wg.Done()
		fmt.Println(string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}
	// Wait for a message to come in
	wg.Wait()
}
