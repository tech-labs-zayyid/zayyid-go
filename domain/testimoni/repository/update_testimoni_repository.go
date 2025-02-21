package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"zayyid-go/domain/testimoni/model"
)

func (t testimoniRepository) UpdateTestimoniRepository(ctx context.Context, request model.Testimoni) (err error) {

	args := []interface{}{}
	buildQuery := []string{}

	if request.Position != "" {
		args = append(args, request.Position)
		buildQuery = append(buildQuery, " position = $1")
	}
	if request.Deskripsi != "" {
		args = append(args, request.Deskripsi)
		buildQuery = append(buildQuery, " deskripsi = $2")
	}
	if request.PhotoUrl != "" {
		args = append(args, request.PhotoUrl)
		buildQuery = append(buildQuery, " photo_url = $3")
	}

	args = append(args, request.IsActive)
	buildQuery = append(buildQuery, " is_active = $4")
	buildQuery = append(buildQuery, " modified_at = NOW()")

	updateQuery := strings.Join(buildQuery, ",")
	args = append(args, request.Id)
	args = append(args, request.UserName)
	query := fmt.Sprintf(`UPDATE testimoni SET %s  WHERE id = ? AND user_name = ? `, updateQuery)

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = errors.New("🔥 [testimoni-prepare-update] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		err = errors.New("🔥 [testimoni-update] | " + err.Error())
		return
	}

	return
}
