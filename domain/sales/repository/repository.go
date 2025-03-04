package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/infrastructure/database"
)

type SalesRepository interface {
	CheckExistsSubdomainSales(ctx context.Context, subdomain string) (exists bool, err error)
	HomeData(ctx context.Context, subdomain string) (resp response.DataHome, err error)

	// product
	CheckExistsProductName(ctx context.Context, productName, salesId string) (exists bool, err error)
	AddProductSales(ctx context.Context, tx *sql.Tx, param request.AddProductReq) (err error)
	GetProductTier(ctx context.Context) (resp response.TierResp, err error)
	GetListProduct(ctx context.Context, filter sharedModel.QueryRequest) (resp map[string]*response.ProductListBySales, err error)
	CountListProduct(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error)
	CheckExistsProductId(ctx context.Context, id string) (exists bool, err error)
	DetailSalesProduct(ctx context.Context, id string) (resp response.ProductDetailResp, err error)
	GetCountDataImageByProductId(ctx context.Context, productId string) (count int, err error)
	UpdateProductSales(ctx context.Context, tx *sql.Tx, param request.UpdateProductSales) (err error)
	ChangeStatusProductSales(ctx context.Context, tx *sql.Tx, param request.UpdateProductSales) (err error)
	GetListProductSalesPublic(ctx context.Context, filter sharedModel.QueryRequest) (resp map[string]*response.ProductListSalesPublic, err error)
	CountListProductSalesPublic(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error)

	// galery
	GetCountDataGalleryBySalesId(ctx context.Context, salesId string) (count int, err error)
	AddGallerySales(ctx context.Context, tx *sql.Tx, param request.AddGalleryParam) (err error)
	GetListDataGallerySales(ctx context.Context, salesId string) (resp response.GalleryResp, err error)
	GetListDataGalleryPublic(ctx context.Context, subdomain string) (resp response.GalleryPublicResp, err error)
	GetDataGallerySales(ctx context.Context, id, salesId string) (resp response.GalleryDataResp, err error)
	UpdateGallerySales(ctx context.Context, req request.UpdateGalleryParam) (err error)

	// banner
	CountBannerSales(ctx context.Context, salesId string) (count int, err error)
	AddBannerSales(ctx context.Context, tx *sql.Tx, param request.BannerReq) (err error)
	GetListBannerSales(ctx context.Context, salesId string) (resp response.BannerListSalesResp, err error)
	GetListBannerPublicSales(ctx context.Context, subdomain string) (resp response.BannerListPublicSalesResp, err error)
	GetBannerSales(ctx context.Context, id, salesId string) (resp response.BannerResp, err error)
	UpdateBannerSales(ctx context.Context, req request.BannerUpdateReq) (err error)

	// testimoni
	AddTestimoniRepository(ctx context.Context, request request.Testimoni) (err error)
	UpdateTestimoniRepository(ctx context.Context, request request.Testimoni) (err error)
	GetTestimoniRepository(ctx context.Context, request request.Testimoni) (response request.Testimoni, err error)
	GetPublicListTestimoniRepository(ctx context.Context, subDomain string, filter request.TestimoniSearch) (response []request.Testimoni, err error)
	GetListTestimoniRepository(ctx context.Context, request request.Testimoni, filter request.TestimoniSearch) (response []request.Testimoni, err error)
	CountListTestimoniRepository(ctx context.Context, request request.Testimoni) (response int, err error)

	//template
	AddTemplateSales(ctx context.Context, param request.AddTemplateReq) (err error)
	GetListTemplateSales(ctx context.Context, salesId string) (resp response.TemplateListSalesResp, err error)
	GetListPublicTemplateSales(ctx context.Context, subdomain string) (resp response.TemplateListPublicResp, err error)
	GetDetailTemplateSales(ctx context.Context, id, salesId string) (resp response.TemplateDetailResp, err error)
	UpdateTemplateSales(ctx context.Context, req request.UpdateTemplateReq) (err error)

	//social media
	AddSocialMediaSales(ctx context.Context, param request.AddSocialMediaReq) (err error)
	GetListSocialMediaSales(ctx context.Context, salesId string) (resp response.SocialMediaListResp, err error)
	GetListPublicSocialMediaSales(ctx context.Context, subdomain string) (resp response.SocialMediaListResp, err error)

	//transaction schema DB
	OpenTransaction(ctx context.Context) (tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx) (rollBack error)
	CommitTransaction(tx *sql.Tx) (commit error)
}

type salesRepository struct {
	database *database.Database
}

func NewSalesRepository(db *database.Database) SalesRepository {
	return &salesRepository{
		database: db,
	}
}
