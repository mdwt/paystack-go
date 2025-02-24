package transactions

import (
	"github.com/mdwt/paystack-go/common"
	"time"
)

// TransactionRequest represents a request to start a transaction.
type TransactionRequest struct {
	CallbackURL       string          `json:"callback_url,omitempty"`
	Reference         string          `json:"reference,omitempty"`
	AuthorizationCode string          `json:"authorization_code,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Amount            float32         `json:"amount,omitempty"`
	Email             string          `json:"email,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	InvoiceLimit      int             `json:"invoice_limit,omitempty"`
	Metadata          common.Metadata `json:"metadata,omitempty"`
	SubAccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int64           `json:"transaction_charge,omitempty"`
	Bearer            string          `json:"bearer,omitempty"`
	Channels          []string        `json:"channels,omitempty"`
}

// TransactionRequest represents a request to start a transaction.
type ChargeAuthorizationRequest struct {
	Amount            int64           `json:"amount" binding:"required"`
	Email             string          `json:"email" binding:"required"`
	AuthorizationCode string          `json:"authorization_code" binding:"required"`
	Reference         string          `json:"reference,omitempty"`
	Currency          string          `json:"currency" binding:"required"`
	Metadata          common.Metadata `json:"metadata,omitempty"`
	Channels          []string        `json:"channels,omitempty"`
	SubAccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int64           `json:"transaction_charge,omitempty"`
	Bearer            string          `json:"bearer,omitempty"`
	Queue             bool            `json:"queue,omitempty"`
}

type InitialiseResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type MetadataUnmarshaler interface {
	UnmarshalMetadata()
}

type ChargeAuthorizationResponse struct {
	Amount          int64       `json:"amount"`
	Currency        string      `json:"currency"`
	TransactionDate time.Time   `json:"transaction_date"`
	Status          string      `json:"status"`
	Reference       string      `json:"reference"`
	Domain          string      `json:"domain"`
	Metadata        interface{} `json:"metadata"`
	GatewayResponse string      `json:"gateway_response"`
	Message         interface{} `json:"message"`
	Channel         string      `json:"channel"`
	IPAddress       interface{} `json:"ip_address"`
	Log             interface{} `json:"log"`
	Fees            int         `json:"fees"`
	Authorization   struct {
		AuthorizationCode string      `json:"authorization_code"`
		Bin               string      `json:"bin"`
		Last4             string      `json:"last4"`
		ExpMonth          string      `json:"exp_month"`
		ExpYear           string      `json:"exp_year"`
		Channel           string      `json:"channel"`
		CardType          string      `json:"card_type"`
		Bank              string      `json:"bank"`
		CountryCode       string      `json:"country_code"`
		Brand             string      `json:"brand"`
		Reusable          bool        `json:"reusable"`
		Signature         string      `json:"signature"`
		AccountName       interface{} `json:"account_name"`
	} `json:"authorization"`
	Customer struct {
		ID           int         `json:"id"`
		FirstName    interface{} `json:"first_name"`
		LastName     interface{} `json:"last_name"`
		Email        string      `json:"email"`
		CustomerCode string      `json:"customer_code"`
		Phone        interface{} `json:"phone"`
		Metadata     struct {
			CustomFields []struct {
				DisplayName  string `json:"display_name"`
				VariableName string `json:"variable_name"`
				Value        string `json:"value"`
			} `json:"custom_fields"`
		} `json:"metadata"`
		RiskAction               string      `json:"risk_action"`
		InternationalFormatPhone interface{} `json:"international_format_phone"`
	} `json:"customer"`
	Plan interface{} `json:"plan"`
	ID   int64       `json:"id"`
}
