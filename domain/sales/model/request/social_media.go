package request

type AddSocialMediaReq struct {
	SalesId         string               `json:"-"`
	PublicAccess    string               `json:"-"`
	DataSocialMedia []SocialMediaListReq `json:"data_social_media"`
}

type SocialMediaListReq struct {
	SocialMediaName string `json:"social_media_name"`
	UserAccount     string `json:"user_account"`
	LinkEmbed       string `json:"link_embed"`
}
