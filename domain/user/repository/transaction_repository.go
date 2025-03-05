package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (r UserRepository) OpenTransaction(ctx context.Context) (tx *sqlx.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}

	tx, _ = r.database.DB.BeginTxx(ctx, &sqlTxOptions)
	return
}

func (r UserRepository) RollbackTransaction(tx *sqlx.Tx) (rollBack error) {
	rollBack = tx.Rollback()

	return
}

func (r UserRepository) CommitTransaction(tx *sqlx.Tx) (commit error) {
	commit = tx.Commit()

	return
}
