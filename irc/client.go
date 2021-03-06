package main

import (
	"fmt"
	"log"
	"net"

	"gopkg.in/irc.v3"
)

func main() {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln(err)
	}

	config := irc.ClientConfig{
		Nick: "i_have_a_nick",
		Pass: "oauth:fegu4qur3ezhe9pc4yqboecl1j4v1n",
		User: "soyPete01",
		Name: "Test McGee",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN soypete01")
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				// Create a handler on all messages.
				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}

	r := irc.NewReader(conn)
	fmt.Println(r.ReadMessage())

}
