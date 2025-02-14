package repository

import (
	"context"
	"database/sql"
	"fmt"

	"zayyid-go/infrastructure/database"
)

type AtomicOperation struct {
	db DBTX
}

type UOWrepository interface {
	ExecTx(ctx context.Context, fn func(*AtomicOperation) error) error
}

type uowRepository struct {
	db *database.Database
}

func NewUOWRepository(db *database.Database) UOWrepository {
	return &uowRepository{
		db: db,
	}
}

// Transaction
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *AtomicOperation {
	return &AtomicOperation{
		db: db,
	}
}

func (u uowRepository) ExecTx(ctx context.Context, fn func(*AtomicOperation) error) error {
	// Start DB Transaction
	sqlOpt := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	}

	tx, err := u.db.BeginTxx(ctx, &sqlOpt)
	if err != nil {
		return err
	}

	newRepo := New(tx)

	err = fn(newRepo)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
