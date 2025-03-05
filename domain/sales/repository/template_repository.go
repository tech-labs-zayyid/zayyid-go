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

func (r salesRepository) AddTemplateSales(ctx context.Context, param request.AddTemplateReq) (err error) {
	query := `INSERT INTO product_marketing.sales_template(id, sales_id, public_access, color_plate_id) VALUES($1, $2, $3, $4)`

	stmt, err := r.database.PrepareContext(ctx, query)
	if err != nil {
		return sharedError.HandleError(err)
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	_, err = stmt.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.SalesId, param.PublicAccess, param.ColorPlate)
	if err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) GetListTemplateSales(ctx context.Context, salesId string) (resp response.TemplateListSalesResp, err error) {
	var (
		data response.DataListTemplate
	)

	query := `SELECT id, color_plate_id FROM product_marketing.sales_template WHERE sales_id = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, salesId)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdTemplate, &data.ColorPlate)
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

func (r salesRepository) GetListPublicTemplateSales(ctx context.Context, subdomain string) (resp response.TemplateListPublicResp, err error) {
	var (
		salesId string
		data    response.DataListTemplate
	)

	query := `SELECT id, sales_id, color_plate_id FROM product_marketing.sales_template WHERE public_access = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.IdTemplate, &salesId, &data.ColorPlate)
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

func (r salesRepository) GetDetailTemplateSales(ctx context.Context, id, salesId string) (resp response.TemplateDetailResp, err error) {
	query := `SELECT id, color_plate_id FROM product_marketing.sales_template WHERE id = $1 AND sales_id = $2`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, id, salesId).Scan(&resp.IdTemplate, &resp.ColorPlate); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) UpdateTemplateSales(ctx context.Context, req request.UpdateTemplateReq) (err error) {
	args := []interface{}{}
	buildQuery := []string{}

	if req.ColorPlate != "" {
		args = append(args, req.ColorPlate)
		buildQuery = append(buildQuery, " color_plate_id = $1")
	}

	buildQuery = append(buildQuery, " updated_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, req.Id)
	args = append(args, req.SalesId)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_template SET %s WHERE id = $%d AND sales_id = $%d `, updateQuery, len(args)-1, len(args))

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

func (r salesRepository) CheckExistsTemplateId(ctx context.Context, id, salesId string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.sales_template 
        	WHERE id = $1 AND sales_id = $2)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, id, salesId).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
