package repository

import (
	"context"
	"fmt"
	"strings"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) GetAgentRepository(ctx context.Context, q model.QueryAgentList) (resp []model.UserRes, err error) {
	
	resp = []model.UserRes{}
	// Base query
	query := `
		SELECT 
			id,
			username,
			name,
			whatsapp_number,
			email,
			role,
			COALESCE(image_url, '') AS image_url,
			COALESCE(referal_code, '') AS referal_code,
			created_at
		FROM 
			product_marketing.users
		WHERE 
			role = 'agent'
	`

	// Parameters for query
	var params []interface{}
	paramIndex := 1

	// Add search filter
	if q.Search != "" {
		query += fmt.Sprintf(" AND (LOWER(username) LIKE $%d OR LOWER(name) LIKE $%d OR LOWER(email) LIKE $%d)", paramIndex, paramIndex+1, paramIndex+2)
		searchParam := "%" + strings.ToLower(q.Search) + "%"
		params = append(params, searchParam, searchParam, searchParam)
		paramIndex += 3
	}

	// Sorting
	sortField := "created_at" // Default sort field
	sortOrder := "DESC"       // Default sort order

	if q.Sort != "" {
		if strings.HasPrefix(q.Sort, "-") {
			sortField = strings.TrimPrefix(q.Sort, "-")
			sortOrder = "DESC"
		} else {
			sortField = q.Sort
			sortOrder = "ASC"
		}

		// Hindari SQL Injection dengan memastikan hanya field yang diperbolehkan
		allowedSortFields := map[string]bool{
			"name":        true,
			"created_at":  true,
			"username":    true,
			"email":       true,
		}

		if !allowedSortFields[sortField] {
			sortField = "created_at"
			sortOrder = "DESC"
		}
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortField, sortOrder)

	// Pagination
	offset := (q.Page - 1) * q.Limit
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
	params = append(params, q.Limit, offset)

	// Prepare query
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return resp, err
	}
	defer stmt.Close()

	// Execute query and fetch data

	err = stmt.SelectContext(ctx, &resp, params...)
	if err != nil {
		return resp, err
	}

	return 

}

