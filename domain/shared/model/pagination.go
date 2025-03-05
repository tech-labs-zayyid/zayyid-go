package model

type Pagination struct {
	Limit     int `json:"limit_per_page"`
	Page      int `json:"current_page"`
	TotalPage int `json:"total_page"`
	TotalRows int `json:"total_rows"`
}

type QueryRequest struct {
	Page               int     `query:"page"`
	Limit              int     `query:"limit"`
	SortBy             string  `query:"sort_by"`
	SortOrder          string  `query:"sort_order"`
	Search             string  `query:"search"`
	Status             string  `query:"status"`
	CityId             string  `query:"city_id"`
	ProvinceId         string  `query:"province_id"`
	SubCategoryProduct string  `json:"sub_category_product"`
	BestProduct        string  `json:"best_product"`
	StatusProduct      string  `json:"status_product"`
	SalesId            string  `json:"sales_id"`
	IsActive           string  `json:"is_active"`
	MinimumPrice       float64 `json:"minimum_price"`
	MaximumPrice       float64 `json:"maximum_price"`
	PublicAccess       string  `json:"public_access"`
}
