// Package main is responsible for calling the actual application.
// The reason for separation of the actual application and main package is
// to allow testing package/different configuration to manipulate without
// depending on only one method i.e. via environment variables
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
