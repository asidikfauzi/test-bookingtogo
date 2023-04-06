package models

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Status  string      `json:"status"`
}
