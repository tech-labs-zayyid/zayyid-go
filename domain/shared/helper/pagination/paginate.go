package pagination

import (
	"context"
	"strconv"
	"zayyid-go/domain/shared/model"
)

func GetPageOrDefault(pageQuery string, defaultValue int) int {
	if pageQuery != "" {
		page, err := strconv.Atoi(pageQuery)
		if err == nil {
			return page
		}
	}
	return defaultValue
}

func GetLimitOrDefault(limitQuery string, defaultValue int) int {
	if limitQuery != "" {
		limit, err := strconv.Atoi(limitQuery)
		if err == nil {
			return limit
		}
	}
	return defaultValue
}

func CalculatePagination(ctx context.Context, limit, totalRows int) (*model.Pagination, error) {

	// Calculate total pages
	totalPages := totalRows / limit
	if totalRows%limit != 0 {
		totalPages++
	}

	return &model.Pagination{
		Limit:     limit,
		TotalRows: totalRows,
		TotalPage: totalPages,
	}, nil
}
