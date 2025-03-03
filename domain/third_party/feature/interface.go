package feature

import (
	"zayyid-go/config"
	"zayyid-go/domain/third_party/repository"
)

type ThirdPartyFeature struct {
	repo   repository.ThirdPartyRepositoryInterface
	config *config.EnvironmentConfig
}

func NewThirdPartyFeature(repo repository.ThirdPartyRepositoryInterface, config *config.EnvironmentConfig) ThirdPartyFeature {
	return ThirdPartyFeature{
		repo:   repo,
		config: config,
	}
}
