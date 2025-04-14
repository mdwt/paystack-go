package transactions

import (
	"context"
	"encoding/json"
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/logger"
)

type Client struct {
	options client.Options
	logger  logger.Logger
}

// New creates a new instance of the transactions client
func New(options client.Options, logger logger.Logger) *Client {
	return &Client{
		options: options,
		logger:  logger,
	}
}

// Initialize initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) Initialize(ctx context.Context, txn *TransactionRequest) (InitialiseResponse, error) {
	jsonData, err := json.Marshal(txn)
	if err != nil {
		return InitialiseResponse{}, err
	}

	resp, err := client.Post[InitialiseResponse](ctx, s.options, "/transaction/initialize", jsonData)
	if err != nil {
		return InitialiseResponse{}, err
	}

	return resp, nil
}

// ChargeAuthorization initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) ChargeAuthorization(ctx context.Context, txn ChargeAuthorizationRequest) (ChargeAuthorizationResponse, error) {
	jsonData, err := json.Marshal(txn)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}

	resp, err := client.Post[ChargeAuthorizationResponse](ctx, s.options, "/transaction/charge_authorization", jsonData)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}

	return resp, nil
}
