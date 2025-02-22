package request

type BannerReq struct {
	Id           string       `json:"-"`
	SalesId      string       `json:"-"`
	PublicAccess string       `json:"-"`
	DataBanner   []DataBanner `json:"data_banner"`
}

type DataBanner struct {
	ImageUrl    string `json:"image_url" validate:"required"`
	Description string `json:"description"`
}
