package response

type TemplateListSalesResp struct {
	SalesId  string             `json:"sales_id"`
	DataList []DataListTemplate `json:"data_list"`
}

type DataListTemplate struct {
	IdTemplate string `json:"id_template"`
	ColorPlate string `json:"color_plate"`
}

type TemplateListPublicResp struct {
	SalesId  string             `json:"sales_id"`
	DataList []DataListTemplate `json:"data_list"`
}

type TemplateDetailResp struct {
	IdTemplate string `json:"id_template"`
	ColorPlate string `json:"color_plate"`
}
