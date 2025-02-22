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
	GetCountDataGalleryBySalesId(ctx context.Context, salesId string) (count int, err error)
	AddGallerySales(ctx context.Context, tx *sql.Tx, param request.AddGalleryParam) (err error)
	GetListDataGallerySales(ctx context.Context, salesId string) (resp response.GalleryResp, err error)
	GetListDataGalleryPublic(ctx context.Context, subdomain string) (resp response.GalleryPublicResp, err error)

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
