package repository

import (
	"database/sql"
	"zayyid-go/infrastructure/database"
)

type UserMenuRepository interface {
	OpenTransaction() (tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx) (rollBack error)
	CommitTransaction(tx *sql.Tx) (commit error)
}

type userMenuRepository struct {
	database *database.Database
}

func NewUserMenuRepository(db *database.Database) UserMenuRepository {
	return &userMenuRepository{
		database: db,
	}
}
