package repository

import (
	"context"
	"zayyid-go/domain/user/model"
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
			sales.users
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
