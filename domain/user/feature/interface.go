package feature

import (
	"context"
	"zayyid-go/domain/user/model"
	"zayyid-go/domain/user/repository"
	"zayyid-go/infrastructure/service/slack"
)

type UserFeature struct {
	repo      repository.UserRepository
	SlackConf slack.SlackNotificationBug
}

func NewUserFeature(repo repository.UserRepository, SlackConf slack.SlackNotificationBug) UserFeature {
	return UserFeature{
		repo:      repo,
		SlackConf: SlackConf,
	}
}

type IUserFeature interface {
	RegisterFeature(ctx context.Context, payload model.RegisterRequest) (resp model.UserRes, err error)
	AuthUserFeature(ctx context.Context, payload model.AuthUserRequest) (resp model.UserRes, err error)
}
