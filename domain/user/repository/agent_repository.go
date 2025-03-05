package repository

import (
	"context"
	"fmt"
	"strings"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) GetAgentRepository(ctx context.Context, q model.QueryAgentList, userId string) (resp []model.UserRes, err error) {
	resp = []model.UserRes{}

	// Base query
	query := `
		SELECT 
			u.id,
			u.username,
			u.name,
			u.whatsapp_number,
			u.email,
			u.role,
			COALESCE(u.image_url, '') AS image_url,
			COALESCE(u.referal_code, '') AS referal_code,
			u.created_at
		FROM 
			product_marketing.users u
		INNER JOIN product_marketing.sales_agent sa ON sa.agent_id = u.id
		WHERE 
			u.role = 'agent' AND 
			sa.sales_id = $1
	`

	// Parameters for query
	var params []interface{}
	params = append(params, userId)
	paramIndex := 2 // Karena $1 sudah digunakan untuk `userId`

	// Add search filter
	if q.Search != "" {
		query += fmt.Sprintf(" AND (LOWER(u.username) LIKE $%d OR LOWER(u.name) LIKE $%d OR LOWER(u.email) LIKE $%d)", paramIndex, paramIndex+1, paramIndex+2)
		searchParam := "%" + strings.ToLower(q.Search) + "%"
		params = append(params, searchParam, searchParam, searchParam)
		paramIndex += 3
	}

	// Sorting
	sortField := "u.created_at" // Default sort field
	sortOrder := "DESC"         // Default sort order

	if q.Sort != "" {
		if strings.HasPrefix(q.Sort, "-") {
			sortField = "u." + strings.TrimPrefix(q.Sort, "-")
			sortOrder = "DESC"
		} else {
			sortField = "u." + q.Sort
			sortOrder = "ASC"
		}

		// Hindari SQL Injection dengan memvalidasi field
		allowedSortFields := map[string]bool{
			"u.name":       true,
			"u.created_at": true,
			"u.username":   true,
			"u.email":      true,
		}

		if !allowedSortFields[sortField] {
			sortField = "u.created_at"
			sortOrder = "DESC"
		}
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortField, sortOrder)

	// Pagination
	if q.Limit > 0 {
		offset := (q.Page - 1) * q.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
		params = append(params, q.Limit, offset)
	}

	// Execute query
	err = r.database.SelectContext(ctx, &resp, query, params...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
