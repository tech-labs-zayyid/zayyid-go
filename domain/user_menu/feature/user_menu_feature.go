package feature

import (
	"context"
	"encoding/json"
	"strings"

	"middleware-cms-api/config"
	errHelper "middleware-cms-api/domain/shared/helper/error"
	atomicRepo "middleware-cms-api/domain/shared/repository"
	"middleware-cms-api/domain/user_menu/helper"
	"middleware-cms-api/domain/user_menu/model"
	"middleware-cms-api/domain/user_menu/repository"
)

type UserMenuFeature struct {
	config     config.EnvironmentConfig
	repo       repository.UserMenuRepository
	atomicRepo atomicRepo.UOWrepository
}

func NewUserMenuFeature(config config.EnvironmentConfig, repo repository.UserMenuRepository, atomicRepo atomicRepo.UOWrepository) *UserMenuFeature {
	return &UserMenuFeature{
		config:     config,
		repo:       repo,
		atomicRepo: atomicRepo,
	}
}

func (f UserMenuFeature) GetList(request model.User) (response model.GetList, err error) {

	if request.Limit == 0 {
		request.Limit = 5
	}

	if request.Page == 0 {
		request.Page = 1
	}

	response, err = f.repo.GetList(request)
	if err != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}

	if len(response.Data) < 1 {
		response.Data = []model.User{}
	}
	response.StatusCode = 200

	return
}

func (f UserMenuFeature) GetDataById(id string) (response *model.UserResponse, err error) {

	resp, err := f.repo.GetDataById(id)
	if err != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}

	var companyPermission []string
	if len(resp.CompanyPermision) > 0 {
		err = json.Unmarshal(resp.CompanyPermision, &companyPermission)
		if err != nil {
			return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
		}
	}

	dataMenu, errDataMenu := f.repo.GetDataUserMenuById(id)
	if errDataMenu != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, errDataMenu.Error())
	}

	var menuId []int
	for _, v := range dataMenu {
		menuId = append(menuId, v.MenuId)
	}

	response = &model.UserResponse{
		Id:               resp.Id,
		Name:             resp.Name,
		Role:             resp.Role,
		CreatedAt:        resp.CreatedAt,
		IsActive:         resp.IsActive,
		CompanyPermision: companyPermission,
		MenuId:           menuId,
	}

	if response.Id == "" {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_DATA_NOT_FOUND, "DATA NOT FOUND")
	}

	response.StatusCode = 200

	return
}

func (f UserMenuFeature) UpsertData(ctx context.Context, request model.User) (response *model.User, err error) {

	isActive := 1
	if !request.IsActive {
		isActive = 0
	}
	request.TempIsActive = isActive

	var companyPermission []string
	err = json.Unmarshal(request.CompanyPermision, &companyPermission)
	if err != nil {
		return
	}

	param, errMarshal := json.Marshal(companyPermission)
	if errMarshal != nil {
		return
	}
	request.CompanyPermision = param

	mainData, errMainData := f.GetDataById(request.Id)
	if errMainData != nil && !strings.Contains(errMainData.Error(), "no rows in result set") {
		err = errMainData
		return
	}

	paramRequest, errCheck := helper.CheckRequest(request)
	if errCheck != nil {
		err = errCheck
		return
	}

	err = f.atomicRepo.ExecTx(ctx, func(ao *atomicRepo.AtomicOperation) error {

		if mainData.Id != "" {
			err = ao.DeleteDataUserAuth(ctx, mainData.Id)
			if err != nil {
				return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
			}
		}

		for _, v := range paramRequest.MenuId {
			dataMenu, errDataMenu := f.repo.GetDataMenuById(v)
			if errDataMenu != nil {
				return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, errDataMenu.Error())
			}

			var permission []string
			err = json.Unmarshal(dataMenu.Permision, &permission)
			if err != nil {
				return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
			}

			paramPermission, errMarshal := json.Marshal(permission)
			if errMarshal != nil {
				return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
			}

			param := model.UserAuth{
				UserId:    paramRequest.Id,
				MenuId:    v,
				Permision: paramPermission,
			}

			err = ao.CreateDataUserAuth(ctx, param)
			if err != nil {
				return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
			}
		}

		if mainData.Id == "" {
			err = ao.CreateDataUser(ctx, paramRequest)
		} else {
			err = ao.UpdateDataUser(ctx, paramRequest)
		}

		if err != nil {
			return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
		}

		return nil
	})

	if err != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}

	response = &request

	return
}

func (f UserMenuFeature) UpsertDataManual(ctx context.Context, request model.User) (response *model.User, err error) {

	isActive := 1
	if !request.IsActive {
		isActive = 0
	}
	request.TempIsActive = isActive

	var companyPermission []string
	var paramCompanyPermission []byte
	if len(request.CompanyPermision) > 0 {
		err = json.Unmarshal(request.CompanyPermision, &companyPermission)
		if err != nil {
			return
		}

		paramCompanyPermission, err = json.Marshal(companyPermission)
		if err != nil {
			return
		}
	}
	request.CompanyPermision = paramCompanyPermission

	mainData, errMainData := f.GetDataById(request.Id)
	if errMainData != nil && !strings.Contains(errMainData.Error(), "no rows in result set") {
		err = errMainData
		return
	}

	paramRequest, errCheck := helper.CheckRequest(request)
	if errCheck != nil {
		err = errCheck
		return
	}

	startTransaction := f.repo.OpenTransaction()
	if mainData.Id != "" {
		err = f.repo.DeleteDataUserAuth(ctx, mainData.Id, startTransaction)
		if err != nil {
			f.repo.RollbackTransaction(startTransaction)
			return
		}
	}

	for _, v := range paramRequest.MenuId {
		dataMenu, errDataMenu := f.repo.GetDataMenuById(v)
		if errDataMenu != nil {
			f.repo.RollbackTransaction(startTransaction)
			return
		}

		var permission []string
		var paramPermission []byte
		if len(dataMenu.Permision) > 0 {
			err = json.Unmarshal(dataMenu.Permision, &permission)
			if err != nil {
				f.repo.RollbackTransaction(startTransaction)
				return
			}

			paramPermission, err = json.Marshal(permission)
			if err != nil {
				f.repo.RollbackTransaction(startTransaction)
				return
			}
		}

		param := model.UserAuth{
			UserId:    paramRequest.Id,
			MenuId:    v,
			Permision: paramPermission,
		}

		err = f.repo.CreateDataUserAuth(ctx, param, startTransaction)
		if err != nil {
			f.repo.RollbackTransaction(startTransaction)
			return
		}
	}

	if mainData.Id == "" {
		err = f.repo.CreateData(ctx, paramRequest, startTransaction)
	} else {
		err = f.repo.UpdateData(ctx, paramRequest, startTransaction)
	}

	if err != nil {
		f.repo.RollbackTransaction(startTransaction)
		return
	}

	f.repo.CommitTransaction(startTransaction)

	response = &request

	return
}

func (f UserMenuFeature) DeleteData(ctx context.Context, id string) (response *model.ParamUserResponse, err error) {

	startTransaction := f.repo.OpenTransaction()
	err = f.repo.DeleteData(ctx, id, startTransaction)
	if err != nil {
		f.repo.RollbackTransaction(startTransaction)
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}
	f.repo.CommitTransaction(startTransaction)

	paramResponse := model.ParamUserResponse{
		Id:         id,
		StatusCode: 200,
	}

	response = &paramResponse

	return
}

func (f UserMenuFeature) GetListMenu() (response *model.GetListMenu, err error) {
	resp, err := f.repo.GetListDataMenu()
	if err != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}

	response = &model.GetListMenu{
		Data:       resp,
		StatusCode: 200,
	}

	return
}

func (f UserMenuFeature) Activate(ctx context.Context, request model.User) (response *model.User, err error) {

	isActive := 1
	if !request.IsActive {
		isActive = 0
	}
	request.TempIsActive = isActive

	err = f.atomicRepo.ExecTx(ctx, func(ao *atomicRepo.AtomicOperation) error {

		err = ao.UpdateDataUser(ctx, request)

		if err != nil {
			return errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
		}

		return nil
	})

	if err != nil {
		return response, errHelper.NewIntegrationError(errHelper.ERROR_INTERNAL_SERVER, err.Error())
	}

	response = &request

	return
}
