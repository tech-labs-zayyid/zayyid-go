package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"zayyid-go/domain/master/model/response"
	"zayyid-go/domain/shared/helper/constant"
	ERR "zayyid-go/domain/shared/helper/error"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/infrastructure/logger"
)

func (r masterRepository) OpenTransaction() (tx *sql.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}
	ctx := context.Background()
	tx, _ = r.database.DB.BeginTx(ctx, &sqlTxOptions)
	return
}

func (r masterRepository) RollbackTransaction(tx *sql.Tx) (rollBack error) {
	rollBack = tx.Rollback()

	return
}

func (r masterRepository) CommitTransaction(tx *sql.Tx) (commit error) {
	commit = tx.Rollback()

	return
}

func (r masterRepository) GetMasterProvince(ctx context.Context, filter sharedModel.QueryRequest) (resp []response.RespProvince, err error) {
	var (
		data     response.RespProvince
		argIndex = 1
	)

	// Calculate offset
	offset := (filter.Page - 1) * filter.Limit

	// Construct the SQL query
	query := `
	SELECT
		id, 
		name,
		is_active,
		created_at,
		updated_at
	FROM 
		sales.master_province mp 
	WHERE 1 = 1`

	args := []interface{}{}

	if filter.Search != "" {
		query += fmt.Sprintf(" AND name LIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.Status != "" {
		isActive, errParse := strconv.ParseBool(filter.Status)
		if errParse != nil {
			err = ERR.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, isActive)
		argIndex++
	}

	if filter.SortBy != "" {
		query += " ORDER BY " + filter.SortBy

		if filter.SortOrder != "" {
			query += " " + filter.SortOrder
		}
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.Limit, offset)

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, args...)
	if err != nil {
		err = ERR.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.IsActive, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			err = ERR.HandleError(err)
			return
		}

		resp = append(resp, data)
	}

	return
}

func (r masterRepository) CountMasterProvince(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error) {
	var (
		argIndex = 1
	)

	// Construct the SQL query
	query := `
	SELECT
		COUNT(id)
	FROM 
		sales.master_province mp 
	WHERE 1 = 1`

	args := []interface{}{}

	if filter.Search != "" {
		query += fmt.Sprintf(" AND name LIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.Status != "" {
		isActive, errParse := strconv.ParseBool(filter.Status)
		if errParse != nil {
			err = ERR.New(http.StatusBadRequest, err.Error(), err)
			return
		}

		query += fmt.Sprintf(" AND is_active = $%d", argIndex)
		args = append(args, isActive)
		argIndex++
	}

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		err = ERR.HandleError(err)
	}

	return
}
