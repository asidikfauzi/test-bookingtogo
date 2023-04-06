package models

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Status  string      `json:"status"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseGetList struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int64       `json:"totalCount"`
}
