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

func (r salesRepository) GetListBannerSales(ctx context.Context, salesId string) (resp response.BannerListSalesResp, err error) {
	var (
		data response.DataListBanner
	)

	query := `SELECT id, image_url, description FROM product_marketing.sales_banner WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, salesId)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdBanner, &data.ImageUrl, &data.Description)
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

func (r salesRepository) GetListBannerPublicSales(ctx context.Context, subdomain string) (resp response.BannerListPublicSalesResp, err error) {
	var (
		salesId string
		data    response.DataListBanner
	)

	query := `SELECT id, sales_id, image_url, description FROM product_marketing.sales_banner WHERE public_access = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdBanner, &salesId, &data.ImageUrl, &data.Description)
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

func (r salesRepository) GetBannerSales(ctx context.Context, id, salesId string) (resp response.BannerResp, err error) {
	query := `SELECT id, image_url, description FROM product_marketing.sales_banner WHERE id = $1 AND sales_id = $2`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, id, salesId).Scan(&resp.IdBanner, &resp.ImageUrl, &resp.Description); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) UpdateBannerSales(ctx context.Context, req request.BannerUpdateReq) (err error) {
	args := []interface{}{}
	buildQuery := []string{}

	if req.ImageUrl != "" {
		args = append(args, req.ImageUrl)
		buildQuery = append(buildQuery, " image_url = $1")
	}
	if req.Description != "" {
		args = append(args, req.Description)
		buildQuery = append(buildQuery, " deskripsi = $2")
	}

	buildQuery = append(buildQuery, " updated_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, req.Id)
	args = append(args, req.SalesId)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_banner SET %s WHERE id = $%d AND sales_id = $%d `, updateQuery, len(args)-1, len(args))

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
