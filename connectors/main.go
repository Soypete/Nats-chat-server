package main

import (
	"encoding/gob"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

func main() {
	gob.Register(&oauth2.Token{})
	redirectURL := "http://localhost:7001/redirect"
	clientID := os.Get("TWITCH_ID")
	clientSecret := os.Get("TWITCH_SECRET")
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:read:email"},
		Endpoint:     twitch.Endpoint,
		RedirectURL:  redirectURL,
	}

	// handleFunc("/redirect", HandleOAuth2)

	// fmt.Println("Started running on http://localhost:7001")
	// fmt.Println(http.ListenAndServe(":7001", nil))
	client := &scopes, http.Client{}
	twitchValues := url.Values{}
	v.Set("client_id", clientID)
	v.Add("client_secret", clientSecret)
	v.Add("code", "")
	v.Add("grant_type", "")
	v.Add("redirect_url", redirectURL)
	req, err := http.NewRequest(http.POST, "https://id.twitch.tv/oauth2/token", nil)
}
