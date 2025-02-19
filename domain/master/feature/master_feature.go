package feature

import (
	"context"
	"zayyid-go/config"
	"zayyid-go/domain/master/model/response"
	"zayyid-go/domain/master/repository"
	paginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
	atomicRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/infrastructure/service/slack"
)

type MasterFeature struct {
	config     config.EnvironmentConfig
	repo       repository.MasterRepository
	atomicRepo atomicRepo.UOWrepository
	SlackConf  slack.SlackNotificationBug
}

func NewMasterFeature(config config.EnvironmentConfig, repo repository.MasterRepository, atomicRepo atomicRepo.UOWrepository, slackConfig slack.SlackNotificationBug) *MasterFeature {
	return &MasterFeature{
		config:     config,
		repo:       repo,
		atomicRepo: atomicRepo,
		SlackConf:  slackConfig,
	}
}

func (f MasterFeature) Ping(ctx context.Context) {
	f.SlackConf.Send("PING :fire:")
}

func (f MasterFeature) MasterProvince(ctx context.Context, paramSearch sharedModel.QueryRequest) (resp []response.RespProvince, pagination *sharedModel.Pagination, err error) {
	resp, err = f.repo.GetMasterProvince(ctx, paramSearch)
	if err != nil {
		return
	}

	count, err := f.repo.CountMasterProvince(ctx, paramSearch)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, paramSearch.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = paramSearch.Page
	return
}

func (f MasterFeature) MasterCity(ctx context.Context, paramSearch sharedModel.QueryRequest) (resp []response.RespCity, pagination *sharedModel.Pagination, err error) {
	resp, err = f.repo.GetMasterCity(ctx, paramSearch)
	if err != nil {
		return
	}

	count, err := f.repo.CountMasterCity(ctx, paramSearch)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, paramSearch.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = paramSearch.Page
	return
}
