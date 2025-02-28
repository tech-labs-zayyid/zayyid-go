package repository

import (
	"context"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) UpdateRepository(ctx context.Context, payload model.RegisterRequest, userId string) (err error) {

	queryUpdate := `
		UPDATE user 
		SET 
			name = $1,
			email = $2,
			password = $3,
			username = $4
		WHERE id = $5
	`

	stmt, err := r.database.PreparexContext(ctx, queryUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, payload.Name, payload.Email, payload.Password, payload.ID)
	if err != nil {
		return err
	}

	return nil

}
