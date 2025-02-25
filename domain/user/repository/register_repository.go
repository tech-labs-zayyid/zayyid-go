package repository

import (
	"context"
	"fmt"
	"zayyid-go/domain/user/model"

	"github.com/google/uuid"
)

func (r UserRepository) RegisterRepository(ctx context.Context, payload model.RegisterRequest) (userId string, err error) {

	// define query for insert 
	query := `
		INSERT INTO 
			sales.users (
				id, 
				username, 
				name, 
				whatsapp_number, 
				email, 
				password, 
				role, 
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
			NOW(), 
			$8
		)
	`

	// Generate id with uuid v7
	uuid, err := uuid.NewV7()
	if err != nil {
		return
	}

	// Preparex 
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		return 
	}
	defer stmt.Close() 

	// Exec context 
	_, err = stmt.ExecContext(ctx, 
		uuid,
		payload.UserName,
		payload.Name,
		payload.WhatsappNumber,
		payload.Email,
		payload.Password,
		payload.Role,
		payload.Email,
	)
	if err != nil {
		return 
	}
	
	userId = uuid.String()

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
			created_at, 
			created_by 
		FROM 
			sales.users 
		WHERE id = $1`

	// Prepare statement
	stmt, err := r.database.PreparexContext(ctx, query)
	if err != nil {
		err = fmt.Errorf("failed to prepare statement: %w", err)
		return 
	}
	defer stmt.Close() 

	// Execute query
	err = stmt.GetContext(ctx, &resp, userId)
	if err != nil {
		fmt.Errorf("failed to get user: %w", err)
		return 
	}

	return 
}