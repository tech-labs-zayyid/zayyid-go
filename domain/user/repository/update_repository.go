package repository

import (
	"context"
	"fmt"
	"zayyid-go/domain/user/model"
)

func (r UserRepository) UpdateRepository(ctx context.Context, payload model.UpdateUser, userId string) error {
	var (
		queryUpdate = "UPDATE product_marketing.users SET"
		args        []interface{}
		argIndex    = 1
	)

	// Dynamic field checks
	if payload.Name != "" {
		queryUpdate += fmt.Sprintf(" name = $%d,", argIndex)
		args = append(args, payload.Name)
		argIndex++
	}
	if payload.Username != "" {
		queryUpdate += fmt.Sprintf(" username = $%d,", argIndex)
		args = append(args, payload.Username)
		argIndex++
	}
	if payload.WhatsappNumber != "" {
		queryUpdate += fmt.Sprintf(" whatsapp_number = $%d,", argIndex)
		args = append(args, payload.WhatsappNumber)
		argIndex++
	}
	if payload.ImageUrl != "" {
		queryUpdate += fmt.Sprintf(" image_url = $%d,", argIndex)
		args = append(args, payload.ImageUrl)
		argIndex++
	}
	if payload.Password != "" {
		queryUpdate += fmt.Sprintf(" password = $%d,", argIndex)
		args = append(args, payload.Password)
		argIndex++
	}

	// Ensure at least one field is being updated
	if len(args) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Always update `updated_at`
	queryUpdate += fmt.Sprintf(" updated_at = NOW() WHERE id = $%d", argIndex)
	args = append(args, userId)

	// Prepare and execute query
	stmt, err := r.database.PreparexContext(ctx, queryUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	return err
}
