package paystack

import (
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/common"
	"github.com/mdwt/paystack-go/logger"
	"github.com/mdwt/paystack-go/refunds"
	"github.com/mdwt/paystack-go/transactions"
)

type Options struct {
	ApiKey    string
	ConnectId string
	Logger    *logger.Logger
}

type PaystackApi struct {
	Transaction *transactions.Client
	Refund      *refunds.Client
}

func NewPaystackApi(options Options) *PaystackApi {
	clientOpts := client.Options{
		ApiKey:    options.ApiKey,
		ConnectId: options.ConnectId,
		BaseUrl:   common.BaseURLV1,
	}

	var log logger.Logger
	if options.Logger != nil {
		log = *options.Logger
	} else {
		log = logger.NewDefaultLogger()
	}

	return &PaystackApi{
		Transaction: transactions.New(clientOpts, log),
		Refund:      refunds.New(clientOpts, log),
	}
}
