package feature

import (
	"context"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedHelper "zayyid-go/domain/shared/helper"
)

func (f SalesFeature) AddTemplateSales(ctx context.Context, param request.AddTemplateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"
	valueCtx.Username = "ekotoyota"

	//validation exists or not sales id in t_gallery

	param.SalesId = valueCtx.SalesId
	param.PublicAccess = valueCtx.Username
	err = f.repo.AddTemplateSales(ctx, param)
	return
}

func (f SalesFeature) GetListTemplateSales(ctx context.Context) (resp response.TemplateListSalesResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetListTemplateSales(ctx, valueCtx.SalesId)
	return
}

func (f SalesFeature) GetListPublicTemplateSales(ctx context.Context, subdomain, referral string) (resp response.TemplateListPublicResp, err error) {
	//validation subdomain

	//validation referral
	if referral != "" {

	}

	resp, err = f.repo.GetListPublicTemplateSales(ctx, subdomain)
	return
}

func (f SalesFeature) GetDetailTemplateSales(ctx context.Context, id string) (resp response.TemplateDetailResp, err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	resp, err = f.repo.GetDetailTemplateSales(ctx, id, valueCtx.SalesId)
	return
}

func (f SalesFeature) UpdateTemplateSales(ctx context.Context, req request.UpdateTemplateReq) (err error) {
	var (
		valueCtx = sharedContext.GetValueContext(ctx)
	)

	err = sharedHelper.Validate(req)
	if err != nil {
		return
	}

	//mocking sales id
	valueCtx.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"

	//validation sales id

	req.SalesId = valueCtx.SalesId
	err = f.repo.UpdateTemplateSales(ctx, req)
	return
}
