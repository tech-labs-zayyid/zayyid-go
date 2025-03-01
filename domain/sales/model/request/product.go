package request

type AddProductReq struct {
	ProductId           string         `json:"-"`
	SalesId             string         `json:"-"`
	PublicAccess        string         `json:"public_access"`
	ProductName         string         `json:"product_name"`
	ProductCategoryName string         `json:"-"`
	ProductSubCategory  string         `json:"product_sub_category"`
	Price               float32        `json:"price" validation:"required"`
	TDP                 float32        `json:"tdp" validation:"required"`
	Installment         float32        `json:"installment" validation:"required"`
	CityId              string         `json:"city_id"`
	BestProduct         bool           `json:"best_product"`
	Description         string         `json:"description"`
	Status              string         `json:"-"`
	Images              []ProductImage `json:"images"`
}

type ProductImage struct {
	ProductId string `json:"-"`
	ImageUrl  string `json:"image_url"`
}
