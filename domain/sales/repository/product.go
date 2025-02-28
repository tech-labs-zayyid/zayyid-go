package repository

import (
	"context"
	"database/sql"
	modelRequest "zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) CheckExistsUserId(ctx context.Context, userId string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.users WHERE id = $1)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, userId).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) AddProductSales(ctx context.Context, tx *sql.Tx, param modelRequest.AddProductReq) (err error) {
	var (
		id = sharedRepo.GenerateUuidAsIdTable().String()
	)

	param.ProductId = sharedRepo.GenerateUuidAsIdTable().String()
	stmtProduct, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product (id, page_category_id, page_category_name, sub_category_product, 
            product_name, price, tdp, installment, best_product, city_id, sales_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductDesc, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_description (id, product_id, description) VALUES ($1, $2, $3)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductStatus, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_status (id, product_id, status_id, status_name) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductImg, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_image (id, product_id, image_url) VALUES ($1, $2, $3)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProduct.ExecContext(ctx, param.ProductId, param.ProductCategoryId, param.ProductCategoryName, param.ProductSubCategory,
		param.ProductName, param.Price, param.TDP, param.Installment, param.BestProduct, param.CityId, param.SalesId); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProductDesc.ExecContext(ctx, id, param.ProductId, param.Description); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProductStatus.ExecContext(ctx, id, param.ProductId, param.StatusId, param.StatusName); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	for _, v := range param.Image {
		id = sharedRepo.GenerateUuidAsIdTable().String()
		if _, err = stmtProductImg.ExecContext(ctx, id, param.ProductId, v.ImageUrl); err != nil {
			err = sharedError.HandleError(err)
			return
		}
	}

	return
}

func (r salesRepository) GetProductTier(ctx context.Context) (resp response.TierResp, err error) {
	query := `SELECT pt.id, pt.tier_name, ptd.feature, ptd.limitation, ptd.length_limitation
			FROM product_marketing.product_tier pt 
			LEFT JOIN product_marketing.product_tier_detail ptd ON pt.id = ptd.tier_id
			WHERE pt.is_active = TRUE AND ptd.is_active = TRUE`

	err = r.database.QueryRowContext(ctx, query).Scan(&resp.Id, &resp.TierName, &resp.Feature, &resp.Limitation, &resp.LengthLimitation)
	if err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
