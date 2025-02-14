package model

type Pagination struct {
	Limit     int64 `json:"limit_per_page"`
	Page      int64 `json:"current_page"`
	TotalPage int64 `json:"total_page"`
	TotalRows int64 `json:"total_rows"`
}
