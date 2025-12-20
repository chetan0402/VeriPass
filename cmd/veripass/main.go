package main

import (
	"os"

	veripass "github.com/chetan0402/veripass/internal"
)

func main() {
	veripass.Run(&veripass.Config{
		DatabaseUrl:    os.Getenv("VERIPASS_DATABASE_URL"),
		OAuthServer:    os.Getenv("OAUTH_SERVER"),
		ClientID:       os.Getenv("CLIENT_ID"),
		ClientSecret:   os.Getenv("CLIENT_SECRET"),
		RedirectionURI: os.Getenv("REDIRECTION_URI"),
	})
}
