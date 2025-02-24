package request

type AddTemplateReq struct {
	SalesId      string `json:"-"`
	PublicAccess string `json:"-"`
	ColorPlate   string `json:"color_plate"`
}

type UpdateTemplateReq struct {
	Id         string `json:"-"`
	SalesId    string `json:"-"`
	ColorPlate string `json:"color_plate" validate:"required"`
}
