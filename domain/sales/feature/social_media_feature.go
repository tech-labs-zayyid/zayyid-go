package feature

import (
	"context"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedError "zayyid-go/domain/shared/helper/error"
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

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"
	valueCtx.Username = "ekotoyota"

	//validation exists or not sales id

	req.SalesId = valueCtx.SalesId
	req.PublicAccess = valueCtx.Username
	err = f.repo.AddSocialMediaSales(ctx, req)
	return
}

func (f salesFeature) GetListSocialMediaSales(ctx context.Context) (resp response.SocialMediaListResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetListSocialMediaSales(ctx, valueCtx.SalesId)
	return
}

func (f salesFeature) GetListSocialMediaPublicSales(ctx context.Context, subdomain, referral string) (resp response.SocialMediaListResp, err error) {
	//validation subdomain

	//validation referral
	if referral != "" {

	}

	resp, err = f.repo.GetListPublicSocialMediaSales(ctx, subdomain)
	return
}
