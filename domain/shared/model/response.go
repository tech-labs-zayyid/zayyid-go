package model

type ResponseWithPagination struct {
	Status     string      `json:"status"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
