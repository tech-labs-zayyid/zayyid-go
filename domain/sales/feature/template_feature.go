package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedHelper "zayyid-go/domain/shared/helper/general"
)

func (f salesFeature) AddTemplateSales(ctx context.Context, param request.AddTemplateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	dataUser, err := f.userRepo.GetDataUserByUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	param.SalesId = valueCtx.UserId
	param.PublicAccess = dataUser.Username
	err = f.repo.AddTemplateSales(ctx, param)
	return
}

func (f salesFeature) GetListTemplateSales(ctx context.Context) (resp response.TemplateListSalesResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListTemplateSales(ctx, valueCtx.SalesId)
	return
}

func (f salesFeature) GetListPublicTemplateSales(ctx context.Context, subdomain string) (resp response.TemplateListPublicResp, err error) {
	exists, err := f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusNotFound, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListPublicTemplateSales(ctx, subdomain)
	return
}

func (f salesFeature) GetDetailTemplateSales(ctx context.Context, id string) (resp response.TemplateDetailResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existTemplateId, err := f.repo.CheckExistsTemplateId(ctx, id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existTemplateId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdTemplateNotFound, errors.New(sharedConstant.ErrIdTemplateNotFound))
		return
	}

	resp, err = f.repo.GetDetailTemplateSales(ctx, id, valueCtx.SalesId)
	return
}

func (f salesFeature) UpdateTemplateSales(ctx context.Context, req request.UpdateTemplateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existTemplateId, err := f.repo.CheckExistsTemplateId(ctx, req.Id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existTemplateId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdTemplateNotFound, errors.New(sharedConstant.ErrIdTemplateNotFound))
		return
	}

	err = sharedHelper.Validate(req)
	if err != nil {
		return
	}

	req.SalesId = valueCtx.UserId
	err = f.repo.UpdateTemplateSales(ctx, req)
	return
}
