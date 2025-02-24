package response

type GalleryResp struct {
	SalesId  string     `json:"sales_id"`
	DataList []DataList `json:"data_list"`
}

type DataList struct {
	IdGallery string `json:"id_gallery"`
	ImageUrl  string `json:"image_url"`
}

type GalleryPublicResp struct {
	SalesId  string     `json:"sales_id"`
	DataList []DataList `json:"data_list"`
}

type GalleryDataResp struct {
	IdGallery string `json:"id_gallery"`
	ImageUrl  string `json:"image_url"`
}
