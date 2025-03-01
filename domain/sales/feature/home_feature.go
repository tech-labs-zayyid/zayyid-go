package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/model/response"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
)

func (f salesFeature) HomeSalesData(ctx context.Context, subdomain, referral string) (resp response.DataHome, err error) {
	exists, err := f.repo.CheckExistsSubdomainSales(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusNotFound, errors.New(sharedConstant.ErrDataNotFound).Error(), errors.New(sharedConstant.ErrDataNotFound))
	}

	//validation referral
	if referral != "" {
		// logic validation referral

		//mocking value
		resp.Agent.PhoneNumber = "0897567474747"
		resp.Agent.Fullname = "dhany"
		resp.Agent.Email = "www@gmail.com"
	}

	resp, err = f.repo.HomeData(ctx, subdomain)
	return
}
