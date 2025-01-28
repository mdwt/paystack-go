package paystack

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// APIError includes the response from the Paystack API and some HTTP request info
type APIError struct {
	Message        string        `json:"message,omitempty"`
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
	URL            string        `json:"url,omitempty"`
	Header         http.Header   `json:"header,omitempty"`
}

// APIError supports the error interface
func (aerr *APIError) Error() string {
	ret, _ := json.Marshal(aerr)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Status  bool                   `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

func newAPIError(resp *resty.Response) *APIError {
	var details ErrorResponse
	_ = json.Unmarshal(resp.Body(), &details)

	return &APIError{
		HTTPStatusCode: resp.StatusCode(),
		Header:         resp.Header(),
		Details:        details,
		URL:            resp.Request.URL,
	}
}
