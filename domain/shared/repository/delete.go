package repository

import (
	"context"
	"errors"
)

func (r AtomicOperation) DeleteDataUserAuth(ctx context.Context, userId string) (err error) {

	query := `
		DELETE FROM
			cms_user_auth
		WHERE
			user_id = ?
	`

	stmt, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		err = errors.New("🔥 [delete] | " + err.Error())
		return
	}

	defer stmt.Close()

	return
}
