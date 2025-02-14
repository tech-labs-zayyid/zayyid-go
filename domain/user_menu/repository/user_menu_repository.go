package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func (r userMenuRepository) OpenTransaction() (tx *sqlx.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}
	ctx := context.Background()
	tx, _ = r.database.DB.BeginTxx(ctx, &sqlTxOptions)
	return
}

func (r userMenuRepository) RollbackTransaction(tx *sqlx.Tx) (rollBack error) {

	rollBack = tx.Rollback()

	return
}

func (r userMenuRepository) CommitTransaction(tx *sqlx.Tx) (commit error) {

	commit = tx.Rollback()

	return
}
