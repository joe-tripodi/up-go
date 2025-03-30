package upgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "https://api.up.com.au/api"

type UpClient struct {
	BearerToken string
	Version     string
	Url         string
	HTTPClient  *http.Client
}

func NewUpClient(accessToken string) (*UpClient, error) {
	upClient := &UpClient{}
	if accessToken == "" {
		return upClient, errors.New("No access token")
	}
	upClient.BearerToken = fmt.Sprintf("Bearer %v", accessToken)
	upClient.Version = "v1"
	upClient.Url = baseUrl + "/" + upClient.Version
	upClient.HTTPClient = &http.Client{}

	return upClient, nil
}

func (c *UpClient) doRequest(method, endpoint string, body any) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.Url, endpoint)
	log.Printf("INFO: %s request to endpoint:%s", method, endpoint)
	log.Printf("INFO: url %s", url)

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API request failed: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
