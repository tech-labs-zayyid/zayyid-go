package repository

import (
	"context"
	"errors"
	"fmt"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/domain/testimoni/model"
	"zayyid-go/infrastructure/logger"
)

func (t testimoniRepository) GetListTestimoniRepository(ctx context.Context, request model.Testimoni, filter model.TestimoniSearch) (response []model.Testimoni, err error) {

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
			testimoni
		WHERE
			1 = 1
			%s %s`, queryCond, queryLimit)
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = errors.New("🔥 [testimoni-prepare-get] | " + err.Error())
		return
	}
	defer stmt.Close()

	err = stmt.SelectContext(ctx, response, args...)
	if err != nil {
		err = errors.New("🔥 [testimoni-get] | " + err.Error())
		return
	}

	return
}

func (t testimoniRepository) CountListTestimoniRepository(ctx context.Context, request model.Testimoni) (response int, err error) {

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
		err = errors.New("🔥 [testimoni-prepare-count] | " + err.Error())
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response, args...)
	if err != nil {
		err = errors.New("🔥 [testimoni-count] | " + err.Error())
		return
	}

	return
}
