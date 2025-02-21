package feature

import (
	"context"
	"fmt"
	"zayyid-go/infrastructure/service/slack"

	"zayyid-go/config"
	atomicRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/domain/testimoni/repository"
)

type TestimoniFeature struct {
	config     config.EnvironmentConfig
	repo       repository.TestimoniRepository
	atomicRepo atomicRepo.UOWrepository
	slackConf  slack.SlackNotificationBug
}

func NewTestimoniFeature(config config.EnvironmentConfig, repo repository.TestimoniRepository, atomicRepo atomicRepo.UOWrepository, slackConfig slack.SlackNotificationBug) *TestimoniFeature {
	return &TestimoniFeature{
		config:     config,
		repo:       repo,
		atomicRepo: atomicRepo,
		slackConf:  slackConfig,
	}
}

func (f TestimoniFeature) Ping(ctx context.Context) {
	err := f.slackConf.Send("Bug from Golang :fire:")
	fmt.Println(err)
}
