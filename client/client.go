package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/mdwt/paystack-go/common"
	"github.com/mdwt/paystack-go/errors"
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

func Post[T any](ctx context.Context, options Options, path string, body []byte) (T, error) {
	var zero T
	req, err := http.NewRequestWithContext(ctx, "POST", options.BaseUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return zero, err
	}
	c := NewApiClient(options)

	resp, err := c.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return zero, errors.NewAPIError(resp)
	}
	var result common.ApiResponse[T]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return zero, err
	}
	
	return result.Data, err
}

func (c *ApiClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Options.ApiKey)
	if c.Options.ConnectId != "" {
		req.Header.Set("X-Connect-Account", c.Options.ConnectId)
	}
	return c.Client.Do(req)
}
