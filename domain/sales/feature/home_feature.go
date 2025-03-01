package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/model/response"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
)

func (f SalesFeature) HomeSalesData(ctx context.Context, subdomain, referral string) (resp response.DataHome, err error) {
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

	}

	resp, err = f.repo.HomeData(ctx, subdomain)

	//mocking value
	resp.Agent.PhoneNumber = "0897567474747"
	resp.Agent.Fullname = "dhany"
	resp.Agent.Email = "www@gmail.com"

	resp.Fullname = "eko eka eke"
	resp.PhoneNumber = "087635342667"
	resp.Email = "eko@mail.com"
	resp.Desc = "Stay Humble, peace love & Gaul"
	resp.UrlImage = "https://res.cloudinary.com/dyj8vcauy/image/upload/v1740813846/hirm-astray_red_frame-1000x1000_nbfhbq.jpg"
	resp.Testimony = append(resp.Testimony, response.TestimonyListHome{
		IdTestimony: "01955073-a98d-7707-954b-540d8c37034b",
		Name:        "Jefri Santuy",
		Description: "waah keyen",
		PhotoUrl:    "https://res.cloudinary.com/dyj8vcauy/image/upload/v1740811698/WhatsApp_Image_2025-02-28_at_17.46.01_ukk40e.jpg",
	})

	resp.Product = append(resp.Product, response.BestProduct{
		IdProduct:          "0195508a-8ec1-7a78-8236-8ab98a4854a7",
		ProductName:        "Lambo Lambe",
		Price:              1.500000000,
		ProductSubCategory: "Sports Car",
		TDP:                700000000,
		Installment:        125000000,
		CityId:             "01951519-c232-775a-b8ef-76c2d0e2fa7e",
		BestProduct:        true,
	})
	return
}
