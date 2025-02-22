package request

type AddGalleryParam struct {
	SalesId      string   `json:"-"`
	PublicAccess string   `json:"public_access"`
	ImageUrl     []string `json:"image_url" validate:"required"`
}
