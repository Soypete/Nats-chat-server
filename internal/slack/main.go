package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SlackEvent struct {
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	APIAppID    string   `json:"api_app_id"`
	Event       Event    `json:"event"`
	Type        string   `json:"type"`
	AuthedTeams []string `json:"authed_teams"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
}
type Event struct {
	Type        string `json:"type"`
	Channel     string `json:"channel"`
	User        string `json:"user"`
	Text        string `json:"text"`
	Ts          string `json:"ts"`
	EventTs     string `json:"event_ts"`
	ChannelType string `json:"channel_type"`
}

func main() {
	http.HandleFunc("/", handleMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	var payload SlackEvent
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
		return
	}
	nanosec, err := strconv.ParseFloat(payload.Event.Ts, 64)
	ts := time.Unix(int64(payload.EventTime), int64(nanosec*10000))
	timestamp := ts.Format(time.Kitchen)
	if err != nil {
		log.Fatal(err)
		return
	}
	message := fmt.Sprintf("%v %s: %s", timestamp, payload.Event.User, payload.Event.Text)
	fmt.Println(message)
	natsbot.Publish(message)
}
