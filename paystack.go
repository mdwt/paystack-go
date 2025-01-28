package paystack

import "github.com/go-resty/resty/v2"

const (
	BaseURLV1 = "https://api.paystack.co"
)

type errorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type PaystackResponse struct {
	Status  int         `json:"status"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type Client struct {
	BaseURL string
	apiKey  string
	client  *resty.Client

	Transaction *TransactionService
}

func NewClient(apiKey string) *Client {
	client := getClient(apiKey)

	return &Client{
		client: client,
		Transaction: &TransactionService{
			client: client,
		},
	}
}

func getClient(apiKey string) *resty.Client {
	client := resty.New()
	client.SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBaseURL(BaseURLV1)
	return client
}
