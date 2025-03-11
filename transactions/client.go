package transactions

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

// Initialize initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) Initialize(ctx context.Context, txn *TransactionRequest) (InitialiseResponse, error) {
	path := fmt.Sprintf("%s/transaction/initialize", s.client.Options.BaseUrl)

	jsonData, err := json.Marshal(txn)
	if err != nil {
		return InitialiseResponse{}, err
	}
	s.logger.Debug("request", "request", jsonData)

	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		return InitialiseResponse{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return InitialiseResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return InitialiseResponse{}, errors.NewAPIError(resp)
	}

	var result common.ApiResponse[InitialiseResponse]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return InitialiseResponse{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return InitialiseResponse{}, err
	}

	return result.Data, err
}

// ChargeAuthorization initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *Client) ChargeAuthorization(ctx context.Context, txn ChargeAuthorizationRequest) (ChargeAuthorizationResponse, error) {
	path := fmt.Sprintf("%s/transaction/charge_authorization", s.client.Options.BaseUrl)
	jsonData, err := json.Marshal(txn)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}
	s.logger.Debug("request", "request", jsonData)
	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return ChargeAuthorizationResponse{}, errors.NewAPIError(resp)
	}

	var result common.ApiResponse[ChargeAuthorizationResponse]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ChargeAuthorizationResponse{}, err
	}

	return result.Data, err
}
