package upgo

import (
	"log"
)

func (up *UpClient) Ping() {
	data, err := up.doRequest("GET", "/util/ping", nil)
	if err != nil {
		log.Fatalf("ERROR: Failed to read response: %v", err)
	}
	log.Printf("INFO: util/ping response: %s", string(data))
}
