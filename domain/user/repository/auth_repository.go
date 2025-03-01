package repository

import (
	"context"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/user/model"
	"zayyid-go/infrastructure/logger"
)

func (r UserRepository) GetByQueryRepository(ctx context.Context, q model.QueryUser) (user model.UserRes, err error) {

	// define query for insert
	query := `
		SELECT 
			id, 
			username, 
			password,
			name, 
			whatsapp_number, 
			email, 
			role, 
			created_at, 
			created_by
		FROM 
			product_marketing.users
		WHERE 	
			email = $1`

	// Preparex
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// Exec context
	err = stmt.GetContext(ctx, &user, q.Email)
	if err != nil {
		return
	}

	return

}

func (r UserRepository) GetDataUserByUserId(ctx context.Context, userId string) (resp model.UserDataResp, err error) {
	query := `SELECT id, username, name, whatsapp_number, email, role FROM product_marketing.users WHERE id = $1`

	// Preparex
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	logger.LogInfo(constant.QUERY, query)
	err = stmt.QueryRowContext(ctx, userId).Scan(&resp.UserId, &resp.Username, &resp.Name, &resp.WhatsappNumber, &resp.Email, &resp.Role)
	if err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
