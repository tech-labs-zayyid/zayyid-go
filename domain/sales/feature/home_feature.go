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
	var (
		exists, existsReferral bool
	)

	exists, err = f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusNotFound, errors.New(sharedConstant.ErrDataNotFound).Error(), errors.New(sharedConstant.ErrDataNotFound))
	}

	//validation referral
	if referral != "" {
		existsReferral, err = f.userRepo.CheckExistsCodeReferal(ctx, referral)
		if err != nil {
			return
		}

		if !existsReferral {
			err = sharedError.New(http.StatusNotFound, errors.New(sharedConstant.ErrReferralCode).Error(), errors.New(sharedConstant.ErrReferralCode))
			return
		}
	}

	resp, product, err := f.repo.HomeData(ctx, subdomain)
	if err != nil {
		return
	}

	for _, v := range product {
		resp.Product = append(resp.Product, *v)
	}

	if existsReferral {
		dataAgent, errAgent := f.userRepo.GetDataAgentByReferralCode(ctx, referral)
		if errAgent != nil {
			err = errAgent
			return
		}

		resp.Agent.Fullname = dataAgent.Name
		resp.Agent.PhoneNumber = dataAgent.WhatsappNumber
		resp.Agent.Email = dataAgent.Email
	}

	return
}
