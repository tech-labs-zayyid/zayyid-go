package feature

import (
	"context"
	"zayyid-go/domain/sales/model/response"
)

func (f SalesFeature) HomeSalesData(ctx context.Context, subdomain, referral string) (resp response.DataHome, err error) {
	//validation subdomain

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
