package common

type ApiResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// ListMeta is pagination metadata for paginated responses from the Paystack API
type ListMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// Metadata is an key-value pairs added to Paystack API requests
type Metadata map[string]interface{}

type MetadataCustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

type errorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type PaystackResponse struct {
	Status  int         `json:"status"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}
