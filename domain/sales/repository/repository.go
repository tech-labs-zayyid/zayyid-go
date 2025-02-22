package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"

	//sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/infrastructure/database"
)

type SalesRepository interface {
	// galery
	GetCountDataGalleryBySalesId(ctx context.Context, salesId string) (count int, err error)
	AddGallerySales(ctx context.Context, tx *sql.Tx, param request.AddGalleryParam) (err error)
	GetListDataGallerySales(ctx context.Context, salesId string) (resp response.GalleryResp, err error)
	GetListDataGalleryPublic(ctx context.Context, subdomain string) (resp response.GalleryPublicResp, err error)
	CountBannerSales(ctx context.Context, salesId string) (count int, err error)
	AddBannerSales(ctx context.Context, tx *sql.Tx, param request.BannerReq) (err error)

	// testimoni
	AddTestimoniRepository(ctx context.Context, request request.Testimoni) (err error)
	UpdateTestimoniRepository(ctx context.Context, request request.Testimoni) (err error)
	GetTestimoniRepository(ctx context.Context, request request.Testimoni) (response request.Testimoni, err error)
	GetListTestimoniRepository(ctx context.Context, request request.Testimoni, filter request.TestimoniSearch) (response []request.Testimoni, err error)
	CountListTestimoniRepository(ctx context.Context, request request.Testimoni) (response int, err error)

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
