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

func (f salesFeature) AddSocialMediaSales(ctx context.Context, req request.AddSocialMediaReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	tx := f.repo.OpenTransaction(ctx)

	defer func() {
		if err != nil {
			errRollback := f.repo.RollbackTransaction(tx)
			if errRollback != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errRollback)
			}
		} else {
			errCommit := f.repo.CommitTransaction(tx)
			if errCommit != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errCommit)
			}
		}
	}()

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

	for _, v := range req.DataSocialMedia {
		if v.LinkEmbed != "" {
			if valid := sharedHelper.IsYouTubeURL(v.LinkEmbed); !valid {
				err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrLinkEmbedNotPermission, errors.New(sharedConstant.ErrLinkEmbedNotPermission))
				return
			}
		}
	}

	req.SalesId = valueCtx.UserId
	req.PublicAccess = dataUser.Username
	err = f.repo.AddSocialMediaSales(ctx, req)
	return
}

func (f salesFeature) GetListSocialMediaSales(ctx context.Context) (resp response.SocialMediaListResp, err error) {
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

	resp, err = f.repo.GetListSocialMediaSales(ctx, valueCtx.UserId)
	return
}

func (f salesFeature) GetListSocialMediaPublicSales(ctx context.Context, subdomain string) (resp response.SocialMediaListResp, err error) {
	exists, err := f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusNotFound, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	resp, err = f.repo.GetListPublicSocialMediaSales(ctx, subdomain)
	return
}

func (f salesFeature) GetDetailSocialMediaSales(ctx context.Context, id string) (resp response.DetailSocialMediaListRes, err error) {
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

	existId, err := f.repo.CheckExistsSocialMediaId(ctx, id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdSocialNotFound, errors.New(sharedConstant.ErrIdSocialNotFound))
		return
	}

	resp, err = f.repo.GetDataSocialMediaSales(ctx, id, valueCtx.UserId)
	return
}

func (f salesFeature) UpdateSocialMediaSales(ctx context.Context, param request.UpdateSocialMediaSales) (err error) {
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

	existId, err := f.repo.CheckExistsSocialMediaId(ctx, param.Id, valueCtx.UserId)
	if err != nil {
		return
	}

	if !existId {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdSocialNotFound, errors.New(sharedConstant.ErrIdSocialNotFound))
		return
	}

	if param.LinkEmbed != "" {
		if valid := sharedHelper.IsYouTubeURL(param.LinkEmbed); !valid {
			err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrLinkEmbedNotPermission, errors.New(sharedConstant.ErrLinkEmbedNotPermission))
			return
		}
	}

	param.SalesId = valueCtx.UserId
	err = f.repo.UpdateSocialMediaSales(ctx, param)
	return
}
