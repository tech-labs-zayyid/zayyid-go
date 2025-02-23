package request

type BannerReq struct {
	SalesId      string       `json:"-"`
	PublicAccess string       `json:"-"`
	DataBanner   []DataBanner `json:"data_banner"`
}

type DataBanner struct {
	ImageUrl    string `json:"image_url" validate:"required"`
	Description string `json:"description"`
}

type BannerUpdateReq struct {
	Id          string `json:"-"`
	SalesId     string `json:"-"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Description string `json:"description"`
}
