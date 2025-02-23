package request

type AddGalleryParam struct {
	SalesId      string   `json:"-"`
	PublicAccess string   `json:"-"`
	ImageUrl     []string `json:"image_url"`
}

type UpdateGalleryParam struct {
	Id       string `json:"-"`
	SalesId  string `json:"-"`
	ImageUrl string `json:"image_url" validate:"required"`
}
