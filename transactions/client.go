package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/common"
	"github.com/mdwt/paystack-go/errors"
	"github.com/mdwt/paystack-go/logger"
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

// Initialize initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) Initialize(ctx context.Context, txn *TransactionRequest) (InitialiseResponse, error) {
	path := fmt.Sprintf("/transaction/initialize")

	jsonData, err := json.Marshal(txn)
	if err != nil {
		return InitialiseResponse{}, err
	}

	s.logger.Debug("request", "request", jsonData)
	resp, err := s.client.R().
		SetContext(ctx).
		SetBody(jsonData).
		Post(path)

	if err != nil {
		return InitialiseResponse{}, err
	}

	if resp.IsError() {
		return InitialiseResponse{}, errors.NewAPIError(resp)
	}

	s.logger.Info("response", "rsp", resp.String())
	var result common.ApiResponse[InitialiseResponse]
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return InitialiseResponse{}, err
	}

	return result.Data, err
}

// ChargeAuthorization initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) ChargeAuthorization(ctx context.Context, txn ChargeAuthorizationRequest) (ChargeAuthorizationResponse, error) {
	path := fmt.Sprintf("/transaction/charge_authorization")
	resp, err := s.client.R().
		SetContext(ctx).
		SetBody(txn).
		Post(path)

	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}

	if resp.IsError() {
		s.logger.Errorf("Error response from Paystack API: %s", resp.String())
		return ChargeAuthorizationResponse{}, errors.NewAPIError(resp)
	}

	s.logger.Debug("response", "rsp", resp.String())
	var result common.ApiResponse[ChargeAuthorizationResponse]
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}

	return result.Data, err
}
