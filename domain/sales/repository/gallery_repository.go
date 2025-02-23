package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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
	query := `INSERT INTO product_marketing.sales_gallery(id, sales_id, public_access, image_url) VALUES($1, $2, $3, $4)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return sharedError.HandleError(err)
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	for _, v := range param.ImageUrl {
		_, err = stmt.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.SalesId, param.PublicAccess, v)
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

	if len(resp.DataList) == 0 {
		err = sharedError.HandleError(sql.ErrNoRows)
		return
	}

	resp.SalesId = salesId
	return
}

func (r salesRepository) GetListDataGalleryPublic(ctx context.Context, subdomain string) (resp response.GalleryPublicResp, err error) {
	var (
		salesId string
		data    response.DataList
	)

	query := `SELECT id, sales_id, image_url FROM product_marketing.sales_gallery WHERE public_access = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdGallery, &salesId, &data.ImageUrl)
		if err != nil {
			err = sharedError.HandleError(err)
			return
		}

		resp.DataList = append(resp.DataList, data)
	}

	if len(resp.DataList) == 0 {
		err = sharedError.HandleError(sql.ErrNoRows)
		return
	}

	resp.SalesId = salesId
	return
}

func (r salesRepository) GetDataGallerySales(ctx context.Context, id, salesId string) (resp response.GalleryDataResp, err error) {
	query := `SELECT id, image_url FROM product_marketing.sales_gallery WHERE id = $1 AND sales_id = $2`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, id, salesId).Scan(&resp.IdGallery, &resp.ImageUrl); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) UpdateGallerySales(ctx context.Context, req request.UpdateGalleryParam) (err error) {
	args := []interface{}{}
	buildQuery := []string{}

	if req.ImageUrl != "" {
		args = append(args, req.ImageUrl)
		buildQuery = append(buildQuery, " image_url = $1")
	}

	buildQuery = append(buildQuery, " updated_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, req.Id)
	args = append(args, req.SalesId)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_gallery SET %s WHERE id = $%d AND sales_id = $%d `, updateQuery, len(args)-1, len(args))

	logger.LogInfo(constant.QUERY, query)
	stmt, err := r.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
