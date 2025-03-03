package repository

import (
	"context"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) RegisterRepository(ctx context.Context, payload model.RegisterRequest, userId string) (err error) {

	// define query for insert
	query := `
		INSERT INTO 
			product_marketing.users (
				id, 
				username, 
				name, 
				whatsapp_number, 
				email, 
				password, 
				role, 
				image_url,
				referal_code,
				created_at, 
				created_by
			)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7, 
			$8,
			$9,
			NOW(), 
			$10
		)
	`

	// Preparex
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// Exec context
	_, err = stmt.ExecContext(ctx,
		userId,
		payload.UserName,
		payload.Name,
		payload.WhatsappNumber,
		payload.Email,
		payload.Password,
		payload.Role,
		payload.ImageUrl,
		payload.ReferalCode,
		payload.Email,
	)
	if err != nil {
		return
	}

	return

}

func (r UserRepository) GetUserById(ctx context.Context, userId string) (resp model.UserRes, err error) {

	// Define query
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
			created_at,  
			created_by 
		FROM 
			product_marketing.users 
		WHERE id = $1`

	// Prepare statement
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return model.UserRes{}, sharedError.HandleError(err)
	}
	defer stmt.Close()

	// Execute query
	err = stmt.GetContext(ctx, &resp, userId)
	if err != nil {
		return model.UserRes{}, sharedError.HandleError(err)
	}

	return

}
