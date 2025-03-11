package client

import (
	"net/http"
)

type ApiClient struct {
	http.Client
	Options
}

type Options struct {
	ApiKey    string
	ConnectId string
	BaseUrl   string
}

func NewApiClient(options Options) *ApiClient {
	client := http.Client{}
	return &ApiClient{
		Client:  client,
		Options: options,
	}
}

func (c *ApiClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Options.ApiKey)
	if c.Options.ConnectId != "" {
		req.Header.Set("X-Connect-Account", c.Options.ConnectId)
	}
	return c.Client.Do(req)
}
