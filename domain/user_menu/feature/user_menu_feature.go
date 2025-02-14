package feature

import (
	"context"
	"fmt"
	"zayyid-go/infrastructure/service/slack"

	"zayyid-go/config"
	atomicRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/user_menu/repository"
)

type UserMenuFeature struct {
	config     config.EnvironmentConfig
	repo       repository.UserMenuRepository
	atomicRepo atomicRepo.UOWrepository
	slackConf  slack.SlackNotificationBug
}

func NewUserMenuFeature(config config.EnvironmentConfig, repo repository.UserMenuRepository, atomicRepo atomicRepo.UOWrepository, slackConfig slack.SlackNotificationBug) *UserMenuFeature {
	return &UserMenuFeature{
		config:     config,
		repo:       repo,
		atomicRepo: atomicRepo,
		slackConf:  slackConfig,
	}
}

func (f UserMenuFeature) Ping(ctx context.Context) {
	err := f.slackConf.Send("Bug from Golang :fire:")
	fmt.Println(err)
}
