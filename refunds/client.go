package refunds

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/common"
	"github.com/mdwt/paystack-go/errors"
	"github.com/mdwt/paystack-go/logger"
	"net/http"
)

type Client struct {
	client *client.ApiClient
	logger logger.Logger
}

// NewClient creates a new instance of the transactions client
func NewClient(apiClient *client.ApiClient, logger logger.Logger) *Client {
	return &Client{
		client: apiClient,
		logger: logger,
	}
}

// Create initiates a refund for a payment.
// For more details see https://paystack.com/docs/api/refund
func (s *Client) Create(ctx context.Context, request *CreateRefundRequest) (CreateRefundResponse, error) {
	path := fmt.Sprintf("%s/refund", s.client.Options.BaseUrl)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return CreateRefundResponse{}, err
	}
	s.logger.Debug("request", "request", jsonData)

	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		return CreateRefundResponse{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return CreateRefundResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return CreateRefundResponse{}, errors.NewAPIError(resp)
	}

	var result common.ApiResponse[CreateRefundResponse]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return CreateRefundResponse{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return CreateRefundResponse{}, err
	}

	return result.Data, err
}
