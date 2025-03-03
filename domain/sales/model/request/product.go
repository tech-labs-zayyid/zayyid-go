package request

type AddProductReq struct {
	ProductId           string         `json:"-"`
	SalesId             string         `json:"-"`
	PublicAccess        string         `json:"-"`
	Slug                string         `json:"-"`
	ProductName         string         `json:"product_name" validation:"required"`
	ProductCategoryName string         `json:"-"`
	ProductSubCategory  string         `json:"product_sub_category"`
	Price               float32        `json:"price" validation:"required"`
	TDP                 float32        `json:"tdp" validation:"required"`
	Installment         float32        `json:"installment" validation:"required"`
	CityId              string         `json:"city_id"`
	BestProduct         bool           `json:"best_product"`
	Description         string         `json:"description" validation:"required"`
	Status              string         `json:"-"`
	Images              []ProductImage `json:"images"`
}

type ProductImage struct {
	ProductId string `json:"-"`
	ImageUrl  string `json:"image_url"`
}

type UpdateProductSales struct {
	ProductId          string               `json:"-"`
	Slug               string               `json:"slug"`
	ProductName        string               `json:"product_name" validation:"required"`
	ProductSubCategory string               `json:"product_sub_category"`
	Price              float32              `json:"price" validation:"required"`
	TDP                float32              `json:"tdp" validation:"required"`
	Installment        float32              `json:"installment" validation:"required"`
	CityId             string               `json:"city_id"`
	BestProduct        bool                 `json:"best_product"`
	IdDescription      string               `json:"id_description"`
	Description        string               `json:"description" validation:"required"`
	Status             string               `json:"status"`
	Images             []ProductImageUpdate `json:"images"`
}

type ProductImageUpdate struct {
	ProductId string `json:"-"`
	ImageUrl  string `json:"image_url"`
	IsActive  bool   `json:"is_active"`
}
