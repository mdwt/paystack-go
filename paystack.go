package paystack

import (
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/transactions"
)

type Options struct {
	ApiKey    string
	ConnectId string
}

type PaystackApi struct {
	Transaction *transactions.Client
}

func NewPaystackApi(options Options) *PaystackApi {
	apiClient := client.NewApiClient(client.Options{
		ApiKey:    options.ApiKey,
		ConnectId: options.ConnectId,
		BaseUrl:   BaseURLV1,
	})

	return &PaystackApi{
		Transaction: transactions.NewClient(apiClient),
	}
}
