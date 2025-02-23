package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/infrastructure/database"
)

type SalesRepository interface {
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
