package repository

import (
	"context"
	"errors"
	"zayyid-go/domain/testimoni/model"
)

func (t testimoniRepository) AddTestimoniRepository(ctx context.Context, request model.Testimoni) (err error) {

	args := []interface{}{
		request.Id,
		request.UserName,
		request.Position,
		request.Deskripsi,
		request.PhotoUrl,
		request.IsActive,
	}

	query := `
		INSERT INTO testimoni (id, user_name, position, deskripsi, photo_url, is_active, created_at)
		VALUES
			($1,$2,$3,$4,$5,$6,NOW())`

	stmt, err := t.database.Preparex(query)
	if err != nil {
		err = errors.New("🔥 [testimoni-prepare-create] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		err = errors.New("🔥 [testimoni-create] | " + err.Error())
		return
	}

	return
}
