package request

type AddGalleryParam struct {
	SalesId  string   `json:"-"`
	ImageUrl []string `json:"image_url"`
}
