package repository

import (
	"zayyid-go/infrastructure/database"

	"github.com/jmoiron/sqlx"
)

type UserMenuRepository interface {
	OpenTransaction() (tx *sqlx.Tx)
	RollbackTransaction(tx *sqlx.Tx) (rollBack error)
	CommitTransaction(tx *sqlx.Tx) (commit error)
}

type userMenuRepository struct {
	database *database.Database
}

func NewUserMenuRepository(db *database.Database) UserMenuRepository {
	return &userMenuRepository{
		database: db,
	}
}
