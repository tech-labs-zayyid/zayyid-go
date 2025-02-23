package response

type BannerListSalesResp struct {
	SalesId  string           `json:"sales_id"`
	DataList []DataListBanner `json:"data_list"`
}

type DataListBanner struct {
	IdBanner    string `json:"id_banner"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type BannerListPublicSalesResp struct {
	SalesId  string           `json:"sales_id"`
	DataList []DataListBanner `json:"data_list"`
}

type BannerResp struct {
	IdBanner    string `json:"id_banner"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}
