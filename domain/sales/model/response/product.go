package response

import "time"

type TierResp struct {
	Id               string `json:"id"`
	TierName         string `json:"tier_name"`
	Feature          string `json:"feature"`
	Limitation       string `json:"limitation"`
	LengthLimitation int    `json:"length_limitation"`
}

type ProductListBySales struct {
	IdProduct          string         `json:"id_product"`
	ProductName        string         `json:"product_name"`
	Price              float32        `json:"price"`
	ProductSubCategory string         `json:"product_sub_category"`
	TDP                float32        `json:"tdp"`
	Installment        float32        `json:"installment"`
	ProvinceId         string         `json:"province_id"`
	ProvinceName       string         `json:"province_name"`
	CityId             string         `json:"city_id"`
	CityName           string         `json:"city_name"`
	BestProduct        bool           `json:"best_product"`
	IdDescription      string         `json:"id_description"`
	Description        string         `json:"description"`
	Status             string         `json:"status"`
	IsActive           bool           `json:"is_active"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	ProductImages      []ProductImage `json:"product_images"`
}

type ProductImage struct {
	ProductImageId string `json:"product_image_id"`
	ImageUrl       string `json:"image_url"`
}

type ProductDetailResp struct {
	IdProduct          string               `json:"id_product"`
	ProductName        string               `json:"product_name"`
	Price              float32              `json:"price"`
	ProductSubCategory string               `json:"product_sub_category"`
	TDP                float32              `json:"tdp"`
	Installment        float32              `json:"installment"`
	ProvinceId         string               `json:"province_id"`
	ProvinceName       string               `json:"province_name"`
	CityId             string               `json:"city_id"`
	CityName           string               `json:"city_name"`
	BestProduct        bool                 `json:"best_product"`
	IdDescription      string               `json:"id_description"`
	Description        string               `json:"description"`
	Status             string               `json:"status"`
	IsActive           bool                 `json:"is_active"`
	ProductImages      []ProductImageDetail `json:"product_images"`
}

type ProductImageDetail struct {
	ProductImageId string `json:"product_image_id"`
	ImageUrl       string `json:"image_url"`
}

type ProductListSalesPublic struct {
	IdProduct          string               `json:"-"`
	Slug               string               `json:"slug"`
	ProductName        string               `json:"product_name"`
	ProductSubCategory string               `json:"product_sub_category"`
	Price              float32              `json:"price"`
	TDP                float32              `json:"tdp"`
	Installment        float32              `json:"installment"`
	ProvinceName       string               `json:"province_name"`
	CityName           string               `json:"city_name"`
	BestProduct        bool                 `json:"best_product"`
	Description        string               `json:"description"`
	Status             string               `json:"status"`
	CreatedAt          time.Time            `json:"-"`
	UpdatedAt          *time.Time           `json:"-"`
	ProductImages      []ProductImagePublic `json:"product_images"`
}

type ProductImagePublic struct {
	ProductImageId string `json:"-"`
	ImageUrl       string `json:"image_url"`
}

type ProductDetailPublicResp struct {
	IdProduct          string                     `json:"id_product"`
	ProductName        string                     `json:"product_name"`
	Price              float32                    `json:"price"`
	ProductSubCategory string                     `json:"product_sub_category"`
	TDP                float32                    `json:"tdp"`
	Installment        float32                    `json:"installment"`
	ProvinceName       string                     `json:"province_name"`
	CityName           string                     `json:"city_name"`
	BestProduct        bool                       `json:"best_product"`
	Description        string                     `json:"description"`
	Status             string                     `json:"status"`
	ProductImages      []ProductImageDetailPublic `json:"product_images"`
}

type ProductImageDetailPublic struct {
	ProductImageId string `json:"product_image_id"`
	ImageUrl       string `json:"image_url"`
}
