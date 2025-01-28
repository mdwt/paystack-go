package paystack

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type TransactionService struct {
	client *resty.Client
}

// TransactionRequest represents a request to start a transaction.
type TransactionRequest struct {
	CallbackURL       string   `json:"callback_url,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	AuthorizationCode string   `json:"authorization_code,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Amount            float32  `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      int      `json:"invoice_limit,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
	SubAccount        string   `json:"subaccount,omitempty"`
	TransactionCharge int      `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
	Channels          []string `json:"channels,omitempty"`
}

type InitialiseResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

// Initialize initiates a transaction process
// For more details see https://paystack.com/docs/api/transaction/#initialize
func (s *TransactionService) Initialize(txn *TransactionRequest) (InitialiseResponse, error) {
	path := fmt.Sprintf("/transaction/initialize")
	resp, err := s.client.R().
		SetBody(txn).
		Post(path)

	if err != nil {
		return InitialiseResponse{}, err
	}

	if resp.IsError() {
		return InitialiseResponse{}, newAPIError(resp)
	}

	var result ApiResponse[InitialiseResponse]
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return InitialiseResponse{}, err
	}

	return result.Data, err
}
