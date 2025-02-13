package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	userMenuModel "zayyid-go/domain/user_menu/model"
)

func (r AtomicOperation) CreateDataUser(ctx context.Context, request userMenuModel.User) (err error) {

	query := `
		INSERT INTO cms_user(
			id,
			name,
			role,
			created_at,
			is_active,
			company_permission
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		request.Id,
		request.Name,
		request.Role,
		sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		request.TempIsActive,
		request.CompanyPermision,
	)

	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}

	return
}

func (r AtomicOperation) CreateDataUserAuth(ctx context.Context, request userMenuModel.UserAuth) (err error) {

	query := `
		INSERT INTO cms_user_auth(
			user_id,
			menu_id,
			permission
		) VALUES (?, ?, ?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		request.UserId,
		request.MenuId,
		request.Permision,
	)

	if err != nil {
		err = errors.New("🔥 [create] | " + err.Error())
		return
	}

	return
}
