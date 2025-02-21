package repository

import (
	"context"
	"errors"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/domain/testimoni/model"
	"zayyid-go/infrastructure/logger"
)

func (t testimoniRepository) GetListTestimoniRepository(ctx context.Context, request model.Testimoni) (response []model.Testimoni, err error) {

	args := []interface{}{}

	queryCond := ""
	if request.UserName != "" {
		args = append(args, request.UserName)
		queryCond += " AND user_name = $1"
	}

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
			testimoni
		WHERE
			1 = 1`
	logger.LogInfo(constant.QUERY, query)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = errors.New("🔥 [testimoni-prepare-get] | " + err.Error())
		return
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, response, args...)
	if err != nil {
		err = errors.New("🔥 [testimoni-get] | " + err.Error())
		return
	}

	return
}
