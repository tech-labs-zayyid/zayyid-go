package repository

import (
	"context"

	"middleware-cms-api/domain/user_menu/model"
	"middleware-cms-api/infrastructure/database"

	"github.com/jmoiron/sqlx"
)

type UserMenuRepository interface {
	GetList(request model.User) (response model.GetList, err error)
	GetDataById(id string) (response model.User, err error)
	GetDataUserMenuById(id string) (response []model.UserAuth, err error)
	GetDataMenuById(id int) (response model.Menu, err error)
	GetListDataMenu() (response []model.Menu, err error)

	CreateData(ctx context.Context, request model.User, tx *sqlx.Tx) (err error)
	UpdateData(ctx context.Context, request model.User, tx *sqlx.Tx) (err error)
	DeleteData(ctx context.Context, id string, tx *sqlx.Tx) (err error)
	CreateDataUserAuth(ctx context.Context, request model.UserAuth, tx *sqlx.Tx) (err error)
	DeleteDataUserAuth(ctx context.Context, userId string, tx *sqlx.Tx) (err error)

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
