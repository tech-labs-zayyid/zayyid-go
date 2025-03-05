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
	GetByQueryRepository(ctx context.Context, q model.QueryUser) (user model.UserRes, err error)
	GetDataUserByUserId(ctx context.Context, userId string) (resp model.UserDataResp, err error)
	UpdateRepository(ctx context.Context, payload model.UpdateUser, userId string) (err error)
	CheckExistsUserId(ctx context.Context, userId string) (exists bool, err error)
	CheckExistsSubdomain(ctx context.Context, subdomain string) (exists bool, err error)
	CheckExistsCodeReferal(ctx context.Context, referal string) (exists bool, err error)
	GetDataAgentByReferralCode(ctx context.Context, referralCode string) (resp model.UserDataResp, err error)
}
