package repository

import (
	"context"
	"fmt"
	"strings"
	modelRequest "zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/infrastructure/logger"
)

func (t salesRepository) AddTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (err error) {

	args := []interface{}{
		request.Id,
		request.UserName,
		request.Position,
		request.Deskripsi,
		request.PhotoUrl,
		request.IsActive,
	}

	query := `
		INSERT INTO product_marketing.sales_testimony (id, user_name, position, deskripsi, photo_url, is_active, created_at)
		VALUES
			($1,$2,$3,$4,$5,$6,NOW())`

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}

func (t salesRepository) UpdateTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (err error) {

	args := []interface{}{}
	buildQuery := []string{}

	if request.Position != "" {
		args = append(args, request.Position)
		buildQuery = append(buildQuery, " position = $1")
	}
	if request.Deskripsi != "" {
		args = append(args, request.Deskripsi)
		buildQuery = append(buildQuery, " deskripsi = $2")
	}
	if request.PhotoUrl != "" {
		args = append(args, request.PhotoUrl)
		buildQuery = append(buildQuery, " photo_url = $3")
	}

	args = append(args, request.IsActive)
	buildQuery = append(buildQuery, " is_active = $4")
	buildQuery = append(buildQuery, " modified_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, request.Id)
	args = append(args, request.UserName)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_testimony SET %s  WHERE id = ? AND user_name = ? `, updateQuery)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}

func (t salesRepository) GetTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (response modelRequest.Testimoni, err error) {

	query := `
		SELECT
			id, 
			user_name, 
			position, 
			deskripsi, 
			photo_url, 
			is_active, 
			created_at,
			modified_at
		FROM
			product_marketing.sales_testimony
		WHERE
			id = $1`
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}

func (t salesRepository) GetListTestimoniRepository(ctx context.Context, request modelRequest.Testimoni, filter modelRequest.TestimoniSearch) (response []modelRequest.Testimoni, err error) {

	var (
		args     []interface{}
		argIndex = 1
	)

	offset := (filter.Page - 1) * filter.Limit

	queryCond := ""
	if request.UserName != "" {
		args = append(args, request.UserName)
		queryCond += fmt.Sprintf(" AND user_name = $%d", argIndex)
		argIndex++
	}

	querySort := ""
	if filter.SortBy != "" {
		querySort += " ORDER BY " + filter.SortBy

		if filter.SortOrder != "" {
			querySort += " " + filter.SortOrder
		}
	}

	args = append(args, filter.Limit, offset)
	queryLimit := fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)

	query := fmt.Sprintf(`
		SELECT
			id, 
			user_name, 
			position, 
			deskripsi, 
			photo_url, 
			is_active, 
			created_at,
			modified_at
		FROM
			product_marketing.sales_testimony
		WHERE
			1 = 1
			%s %s`, queryCond, queryLimit)
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	err = stmt.SelectContext(ctx, response, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}

func (t salesRepository) CountListTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (response int, err error) {

	var (
		args     []interface{}
		argIndex = 1
	)

	queryCond := ""
	if request.UserName != "" {
		args = append(args, request.UserName)
		queryCond += fmt.Sprintf(" AND user_name = $%d", argIndex)
		argIndex++
	}

	query := fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			testimoni
		WHERE
			1 = 1
			%s`, queryCond)
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}
