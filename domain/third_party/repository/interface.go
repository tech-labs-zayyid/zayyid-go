package repository

import (
	"context"
	"zayyid-go/domain/third_party/model"
	"zayyid-go/infrastructure/database"
)

type ThirdPartyRepositoryInterface interface {
	AddSalesPaymentRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (err error)
	UpdateSalesPaymentRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (err error)
	GetSalesPaymentRepository(ctx context.Context, request model.FrontendNotificationBodyReq) (response model.SalesPaymentResp, err error)
}
type thirdPartyRepository struct {
	database *database.Database
}

func NewThirdPartyRepository(db *database.Database) ThirdPartyRepositoryInterface {
	return thirdPartyRepository{
		database: db,
	}
}
