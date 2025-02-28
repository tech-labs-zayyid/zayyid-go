package request

import "zayyid-go/domain/sales/helper"

type AddProductReq struct {
	ProductId           string                  `json:"-"`
	SalesId             string                  `json:"-"`
	ProductName         string                  `json:"product_name"`
	ProductCategoryId   helper.PageCategoryType `json:"-"`
	ProductCategoryName string                  `json:"-"`
	ProductSubCategory  string                  `json:"product_sub_category"`
	Price               float32                 `json:"price" validation:"required"`
	TDP                 float32                 `json:"tdp" validation:"required"`
	Installment         float32                 `json:"installment" validation:"required"`
	CityId              string                  `json:"city_id"`
	BestProduct         bool                    `json:"best_product"`
	Description         string                  `json:"description"`
	StatusId            helper.StatusProduct    `json:"-"`
	StatusName          string                  `json:"-"`
	Image               []ProductImage          `json:"image"`
}

type ProductImage struct {
	ProductId string `json:"-"`
	ImageUrl  string `json:"image_url"`
}
