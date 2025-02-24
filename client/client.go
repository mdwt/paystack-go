package client

import "github.com/go-resty/resty/v2"

type ApiClient struct {
	*resty.Client
}

type Options struct {
	ApiKey    string
	ConnectId string
	BaseUrl   string
}

func NewApiClient(options Options) *ApiClient {
	client := resty.New()
	client.SetAuthToken(options.ApiKey).
		SetBaseURL(options.BaseUrl).
		SetHeader("Content-Type", "application/json")

	if options.ConnectId != "" {
		client.SetHeader("X-Connect-Account", options.ConnectId)
	}

	return &ApiClient{
		Client: client,
	}
}
