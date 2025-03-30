package main

import (
	"log"
	"os"

	"github.com/joe-tripodi/up-go"
)

func main() {
	accessToken := os.Getenv("UP_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatalf("ERROR: UP_ACCESS_TOKEN must be set")
	}
	client, err := upclient.NewUpClient(accessToken)

	if err != nil {
		log.Fatalf("ERROR: Unable to create UP Client: %s", err)
	}

	client.Ping()
	client.GetAccounts()
}
