package feature

import (
	"context"
	"fmt"
	"zayyid-go/infrastructure/service/slack"

	"zayyid-go/config"
	"zayyid-go/domain/testimoni/repository"
)

type TestimoniFeature struct {
	config    config.EnvironmentConfig
	repo      repository.TestimoniRepository
	SlackConf slack.SlackNotificationBug
}

func NewTestimoniFeature(config config.EnvironmentConfig, repo repository.TestimoniRepository, slackConfig slack.SlackNotificationBug) *TestimoniFeature {
	return &TestimoniFeature{
		config:    config,
		repo:      repo,
		SlackConf: slackConfig,
	}
}

func (f TestimoniFeature) Ping(ctx context.Context) {
	err := f.SlackConf.Send("Bug from Golang :fire:")
	fmt.Println(err)
}
