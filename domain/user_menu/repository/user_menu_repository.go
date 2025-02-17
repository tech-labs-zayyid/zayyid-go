package repository

import (
	"context"
	"database/sql"
)

func (r userMenuRepository) OpenTransaction() (tx *sql.Tx) {
	sqlTxOptions := sql.TxOptions{
		Isolation: sql.LevelDefault,
	}
	ctx := context.Background()
	tx, _ = r.database.DB.BeginTx(ctx, &sqlTxOptions)
	return
}

func (r userMenuRepository) RollbackTransaction(tx *sql.Tx) (rollBack error) {

	rollBack = tx.Rollback()

	return
}

func (r userMenuRepository) CommitTransaction(tx *sql.Tx) (commit error) {

	commit = tx.Rollback()

	return
}
