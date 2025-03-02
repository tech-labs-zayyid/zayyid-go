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

func (r salesRepository) AddTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (err error) {

	args := []interface{}{
		request.Id,
		request.PublicAccess,
		request.FullName,
		request.Description,
		request.PhotoUrl,
		request.IsActive,
	}

	query := `
		INSERT INTO product_marketing.sales_testimony (id, public_access, fullname, description, photo_url, is_active)
		VALUES
			($1,$2,$3,$4,$5,$6)`

	stmt, err := r.database.Preparex(query)
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

func (r salesRepository) UpdateTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (err error) {

	args := []interface{}{}
	buildQuery := []string{}

	if request.FullName != "" {
		args = append(args, request.FullName)
		buildQuery = append(buildQuery, " fullname = $1")
	}
	if request.Description != "" {
		args = append(args, request.Description)
		buildQuery = append(buildQuery, " description = $2")
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
	args = append(args, request.PublicAccess)
	query := fmt.Sprintf(`UPDATE product_marketing.sales_testimony SET %s  WHERE id = ? AND public_access = ? `, updateQuery)

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
		return
	}

	return
}

func (r salesRepository) GetTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (response modelRequest.Testimoni, err error) {

	query := `
		SELECT
			id, 
			oublic_access, 
			fullname, 
			description, 
			photo_url, 
			is_active, 
			created_at,
			modified_at
		FROM
			product_marketing.sales_testimony
		WHERE
			id = $1`
	logger.LogInfo(constant.QUERY, query)

	stmt, err := r.database.Preparex(query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response, request.Id)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	return
}

func (r salesRepository) GetPublicListTestimoniRepository(ctx context.Context, request string, filter modelRequest.TestimoniSearch) (response []modelRequest.Testimoni, err error) {

	var (
		args     []interface{}
		argIndex = 1
	)

	offset := (filter.Page - 1) * filter.Limit

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
			public_access, 
			fullname, 
			description, 
			photo_url, 
			is_active, 
			created_at,
			modified_at
		FROM
			product_marketing.sales_testimony
		WHERE
			1 = 1
			AND is_active = TRUE
			AND public_access = $1
			%s`, queryLimit)
	logger.LogInfo(constant.QUERY, query)

	stmt, err := r.database.Preparex(query)
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

func (r salesRepository) GetListTestimoniRepository(ctx context.Context, request modelRequest.Testimoni, filter modelRequest.TestimoniSearch) (response []modelRequest.Testimoni, err error) {

	var (
		args     []interface{}
		argIndex = 1
	)

	offset := (filter.Page - 1) * filter.Limit

	queryCond := ""
	if request.PublicAccess != "" {
		args = append(args, request.PublicAccess)
		queryCond += fmt.Sprintf(" AND public_access = $%d", argIndex)
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
			fullname, 
			public_access, 
			description, 
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

	stmt, err := r.database.Preparex(query)
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

func (r salesRepository) CountListTestimoniRepository(ctx context.Context, request modelRequest.Testimoni) (response int, err error) {

	var (
		args     []interface{}
		argIndex = 1
	)

	queryCond := ""
	if request.PublicAccess != "" {
		args = append(args, request.PublicAccess)
		queryCond += fmt.Sprintf(" AND public_access = $%d", argIndex)
		argIndex++
	}

	query := fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			product_marketing.sales_testimony
		WHERE
			1 = 1
			%s`, queryCond)
	logger.LogInfo(constant.QUERY, query)

	stmt, err := r.database.Preparex(query)
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
