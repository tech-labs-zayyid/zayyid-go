package response

type DataHome struct {
	Fullname    string                    `json:"fullname"`
	PhoneNumber string                    `json:"phone_number"`
	Email       string                    `json:"email"`
	Desc        string                    `json:"desc"`
	Gallery     []GalleryDataHomeResp     `json:"gallery"`
	Banner      []BannerHomeResp          `json:"banner"`
	SocialMedia []DataListSocialMediaHome `json:"social_media"`
	Template    []DataListTemplateHome    `json:"template"`
	Testimony   []TestimonyListHome       `json:"testimony"`
	Product     []BestProduct             `json:"product"`
	Agent       DataAgent                 `json:"agent"`
}

type GalleryDataHomeResp struct {
	IdGallery string `json:"id_gallery"`
	ImageUrl  string `json:"image_url"`
}

type BannerHomeResp struct {
	IdBanner    string `json:"id_banner"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type DataListSocialMediaHome struct {
	IdSocialMedia   string `json:"id_social_media"`
	SocialMediaName string `json:"social_media_name"`
	UserAccount     string `json:"user_account"`
	LinkEmbed       string `json:"link_embed"`
}

type DataListTemplateHome struct {
	IdTemplate string `json:"id_template"`
	ColorPlate string `json:"color_plate"`
}

type TestimonyListHome struct {
	IdTestimony string `json:"id_testimony"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoUrl    string `json:"photo_url"`
}

type BestProduct struct {
	IdProduct   string  `json:"id_product"`
	ProductName string  `json:"product_name"`
	Price       float32 `json:"price"`
}

type DataAgent struct {
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
