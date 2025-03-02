package repository

import (
	"context"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) UpdateRepository(ctx context.Context, payload model.UpdateUser, userId string) (err error) {

	queryUpdate := `
		UPDATE product_marketing.users 
		SET 
			name = $1,
			email = $2,
			updated_at = NOW()
		WHERE id = $3
	`

	// prepare context
	stmt, err := r.database.PreparexContext(ctx, queryUpdate)
	if err != nil {
		return
	}

	// execute context
	_, err = stmt.ExecContext(ctx,
		payload.Name,
		payload.Username,
		payload.ImageUrl,
		payload.WhatsappNumber,
		payload.Password,
	)
	if err != nil {
		return err
	}

	return

}
