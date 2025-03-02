package feature

import (
	"zayyid-go/config"
	"zayyid-go/domain/third_party/repository"
	"zayyid-go/infrastructure/service/slack"
)

type ThirdPartyFeature struct {
	repo      repository.ThirdPartyRepositoryInterface
	SlackConf slack.SlackNotificationBug
	config    *config.EnvironmentConfig
}

func NewThirdPartyFeature(repo repository.ThirdPartyRepositoryInterface, SlackConf slack.SlackNotificationBug, config *config.EnvironmentConfig) ThirdPartyFeature {
	return ThirdPartyFeature{
		repo:      repo,
		SlackConf: SlackConf,
		config:    config,
	}
}
