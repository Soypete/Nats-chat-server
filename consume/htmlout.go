package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

type Message struct {
	Time string
	Text string
}
type Data struct {
	Messages []Message // TODO: make this fix length array that holds 15 newest message
}

// TODO: Serve this up on nginx

// var data Data

func Handler(m *stan.Msg) {
	msg := Message{
		Time: time.Now().Format(time.Kitchen),
		Text: string(m.Data),
	}
	fmt.Println(msg)
}

// how do I refresh?
func main() {
	nc, err := nats.Connect("0.0.0.0:4222")
	nc.Flush()
	defer nc.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nc.ConnectedServerId())
	sc, err := stan.Connect("NDEG6G4TGMFG2MFF5HWIAMSTEOYMHYBJ4VLBBMCYCNYKBRON6SHN5WKO", "", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}

	startOpt := stan.StartAt(pb.StartPosition_NewOnly)

	// Subscribe
	sc.QueueSubscribe("CHATS", "", Handler, startOpt)

	// tmpl := template.Must(template.ParseFiles("resources/index.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// tmpl.Execute(w, data)
	// })
	// http.ListenAndServe(":80", nil)
}
