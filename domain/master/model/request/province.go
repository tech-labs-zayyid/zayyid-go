package request

type FilterProvinceParam struct {
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	SortBy     string `query:"sort_by"`
	SortOrder  string `query:"sort_order"`
	Search     string `query:"search"`
	Status     bool   `query:"status"`
	CityId     string `query:"city_id"`
	ProvinceId string `query:"province_id"`
}
