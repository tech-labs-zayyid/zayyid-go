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

type ProductListPublic struct {
	Page               int    `json:"page" query:"page" default:"1"`
	Limit              int    `json:"limit" query:"limit" default:"20"`
	SortBy             string `json:"sort_by" query:"sort_by"`
	SortOrder          string `json:"sort_order" query:"sort_order"`
	Search             string `json:"search" query:"search"`
	SubCategoryProduct string `json:"sub_category_product" query:"sub_category_product"`
	BestProduct        string `json:"best_product" query:"best_product"`
	StatusProduct      string `json:"status_product" query:"status_product"`
	SalesId            string `json:"sales_id" query:"sales_id"`
	MinimumPrice       string `json:"minimum_price" query:"minimum_price"`
	MaximumPrice       string `json:"maximum_price" query:"maximum_price"`
}
