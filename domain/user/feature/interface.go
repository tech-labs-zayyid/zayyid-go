package feature

import (
	"context"
	"zayyid-go/domain/user/model"
	"zayyid-go/domain/user/repository"
)

type UserFeature struct {
	repo      repository.UserRepository
}

func NewUserFeature( repo repository.UserRepository) UserFeature {
	return UserFeature{
		repo:      repo,
	}
}

type IUserFeature interface {
	RegisterFeature(ctx context.Context, payload model.RegisterRequest) (resp model.UserRes, err error)
}