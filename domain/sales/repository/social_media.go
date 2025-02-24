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

func (r salesRepository) AddSocialMediaSales(ctx context.Context, param request.AddSocialMediaReq) (err error) {
	query := `INSERT INTO product_marketing.sales_social_media(id, sales_id, public_access, social_media_name, user_account, link_embed) VALUES($1, $2, $3, $4, $5, $6)`

	stmt, err := r.database.PrepareContext(ctx, query)
	if err != nil {
		return sharedError.HandleError(err)
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	for _, v := range param.DataSocialMedia {
		_, err = stmt.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.SalesId, param.PublicAccess, v.SocialMediaName, v.UserAccount, v.LinkEmbed)
		if err != nil {
			err = sharedError.HandleError(err)
		}
	}

	return
}

func (r salesRepository) GetListSocialMediaSales(ctx context.Context, salesId string) (resp response.SocialMediaListResp, err error) {
	var (
		data response.DataListSocialMedia
	)

	query := `SELECT id, social_media_name, user_account, link_embed FROM product_marketing.sales_social_media WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, salesId)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdSocialMedia, &data.SocialMediaName, &data.UserAccount, &data.LinkEmbed)
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

func (r salesRepository) GetListPublicSocialMediaSales(ctx context.Context, subdomain string) (resp response.SocialMediaListResp, err error) {
	var (
		salesId string
		data    response.DataListSocialMedia
	)

	query := `SELECT id, sales_id, social_media_name, user_account, link_embed FROM product_marketing.sales_social_media WHERE public_access = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdSocialMedia, &salesId, &data.SocialMediaName, &data.UserAccount, &data.LinkEmbed)
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
