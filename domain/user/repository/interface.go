package repository

import (
	"context"
	"zayyid-go/domain/user/model"
	"zayyid-go/infrastructure/database"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{
		database: db,
	}
}

type IUserRepository interface {
	RegisterRepository(ctx context.Context, payload model.RegisterRequest, userId string) (err error)
	GetUserById(ctx context.Context, userId string) (resp model.UserRes, err error)
}
