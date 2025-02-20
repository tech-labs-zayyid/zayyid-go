package repository

import (
	"context"
	"database/sql"
)

func (r masterRepository) OpenTransaction(ctx context.Context) (tx *sql.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}

	tx, _ = r.database.DB.BeginTx(ctx, &sqlTxOptions)
	return
}

func (r masterRepository) RollbackTransaction(tx *sql.Tx) (rollBack error) {
	rollBack = tx.Rollback()

	return
}

func (r masterRepository) CommitTransaction(tx *sql.Tx) (commit error) {
	commit = tx.Commit()

	return
}
