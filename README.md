# WIP:Nats-chat-server

[Nats](https://nats.io/) is probably not needed for all this fun goodness, but I want to experiment with it, so why not.
This is a Chat message subscription service meant to copy the functionality of the [Streamlabs-twitch chat bot](https://streamlabs.com/). The key difference that user input will come from a variety of sources before going to the publish ui. Probably doesn't need Nats but we want it. 

## TODO:
* ~Test out Nats server with connection and publishing~
* ~Get user input to nats server~
* add twitter/twitch/discord/slack api consumption
	Useful URLS:
	* [twitch oauth](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth)
	* [golang](https://golang.org/pkg/net/url/)
	* [twitch auth example in Go](https://github.com/twitchdev/authentication-go-sample/blob/3ee8389a80a6d41bb3d70466797882e523884875/oauth-authorization-code/main.go)
* add moderator interface
* publish chat to web address.


*NOTES*
here is a [link](https://synadia.com/ngs/pricing) for hosted nats jetstream is you want to use the free tier in the future... miriah. 
