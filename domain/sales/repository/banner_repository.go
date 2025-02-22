package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) CountBannerSales(ctx context.Context, salesId string) (count int, err error) {
	query := `
	SELECT
		COUNT(id)
	FROM 
		product_marketing.sales_banner
	WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, salesId).Scan(&count); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) AddBannerSales(ctx context.Context, tx *sql.Tx, param request.BannerReq) (err error) {
	query := `INSERT INTO product_marketing.sales_banner(id, sales_id, public_access, image_url, description) VALUES($1, $2, $3, $4, $5)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return sharedError.HandleError(err)
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	for _, v := range param.DataBanner {
		_, err = stmt.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.SalesId, param.PublicAccess, v.ImageUrl, v.Description)
		if err != nil {
			return sharedError.HandleError(err)
		}
	}

	return
}
