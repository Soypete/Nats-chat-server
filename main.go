package main

import (
	"math/rand"
	"time"

	n "github.com/Soypete/Nats-chat-server/internal/nats-publisher"
)

func RunSpam(nc *n.NatsClient) {
	for {
		// generate random sentences
		sentence := getText()
		// publish sentence
		nc.Publish(sentence)
		// wait 30 sec
		time.Sleep(30)
	}
}

func main() {
	nc, err := n.Setup()
	if err != nil {
		panic(err)
	}
	defer nc.Connection.Close()
	RunSpam(nc)
}

func getText() string {
	rand.Seed(time.Now().Unix())
	answers := []string{
		"Follow soypete01 on twitch",
		"Check out my meetups",
		"Follow Me on Twitter",
		"Do you want to see more of my dogs?",
		"Say hi in chat",
		"Do you want me to work on the cloud technologies",
	}
	return answers[rand.Intn(len(answers))]
}
