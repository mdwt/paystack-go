package errors

import (
	"encoding/json"
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
func (e APIError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Status  bool                   `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

func NewAPIError(resp *http.Response) *APIError {
	var details ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&details)
	if err != nil {
		details = ErrorResponse{
			Status:  false,
			Message: "Failed to decode error response",
		}
	}

	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		Header:         resp.Header,
		Details:        details,
		URL:            resp.Request.URL.String(),
	}
}
