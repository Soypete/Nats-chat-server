package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func main() {
	gob.Register(&oauth2.Token{})
	redirectURL := "http://localhost:7001/redirect"
	clientID := os.Getenv("TWITCH_ID")
	clientSecret := os.Getenv("TWITCH_SECRET")
	// oauth2Config := &oauth2.Config{
	// ClientID:     clientID,
	// ClientSecret: clientSecret,
	// Scopes:       []string{"user:read:email"},
	// Endpoint:     twitch.Endpoint,
	// RedirectURL:  redirectURL,
	// }

	// handleFunc("/redirect", HandleOAuth2)

	// fmt.Println("Started running on http://localhost:7001")
	// fmt.Println(http.ListenAndServe(":7001", nil))
	client := http.Client{}
	// twitchValues := url.Values{}
	// twitchValues.Set("client_id", clientID)
	// twitchValues.Add("client_secret", clientSecret)
	// twitchValues.Add("code", "")
	// twitchValues.Add("grant_type", "")
	// twitchValues.Add("redirect_url", redirectURL)

	//define scope
	//moderation:read
	//chat:read
	//chat:edit

	// read the twitch chat and write to twitch chat

	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", nil)
	resp, err := client.Do(req)
	fmt.Println(resp, err)
}
