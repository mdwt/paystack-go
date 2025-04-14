package refunds

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

// Create initiates a refund for a payment.
// For more details see https://paystack.com/docs/api/refund
func (s *Client) Create(ctx context.Context, request *CreateRefundRequest) (CreateRefundResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return CreateRefundResponse{}, err
	}

	resp, err := client.Post[CreateRefundResponse](ctx, s.options, "/refund", jsonData)
	if err != nil {
		return CreateRefundResponse{}, err
	}

	return resp, err
}
