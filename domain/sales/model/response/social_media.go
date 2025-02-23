package response

type SocialMediaListResp struct {
	SalesId  string                `json:"sales_id"`
	DataList []DataListSocialMedia `json:"data_list"`
}

type DataListSocialMedia struct {
	IdSocialMedia   string `json:"id_social_media"`
	SocialMediaName string `json:"social_media_name"`
	UserAccount     string `json:"user_account"`
	LinkEmbed       string `json:"link_embed"`
}
