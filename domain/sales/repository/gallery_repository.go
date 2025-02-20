package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) GetCountDataGalleryBySalesId(ctx context.Context, salesId string) (count int, err error) {
	query := `
	SELECT
		COUNT(id)
	FROM 
		product_marketing.sales_gallery 
	WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, salesId).Scan(&count); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) AddGallerySales(ctx context.Context, tx *sql.Tx, param request.AddGalleryParam) (err error) {
	query := `INSERT INTO product_marketing.sales_gallery(id, sales_id, image_url) VALUES($1, $2, $3)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return sharedError.HandleError(err)
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	for _, v := range param.ImageUrl {
		_, err = stmt.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.SalesId, v)
		if err != nil {
			return sharedError.HandleError(err)
		}
	}

	return
}

func (r salesRepository) GetListDataGallerySales(ctx context.Context, salesId string) (resp response.GalleryResp, err error) {
	var (
		data response.DataList
	)

	query := `SELECT id, image_url FROM product_marketing.sales_gallery WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, salesId)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdGallery, &data.ImageUrl)
		if err != nil {
			err = sharedError.HandleError(err)
			return
		}

		resp.DataList = append(resp.DataList, data)
	}

	resp.SalesId = salesId
	return
}
