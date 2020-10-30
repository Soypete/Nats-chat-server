package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("0.0.0.0:4222", nats.FlusherTimeout(time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// Do something with the connection
	nc.Publish("Test-message", []byte("spam"))
	// ND3UB7RH2YQFYNULXAPWXHCFMDDVFUUTB2YYBLOZNK3RWTVZP75TKMOC

}
