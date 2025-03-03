package feature

import (
	"context"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/sales/repository"
	sharedModel "zayyid-go/domain/shared/model"
	repoUser "zayyid-go/domain/user/repository"
)

type SalesFeature interface {
	HomeSalesData(ctx context.Context, subdomain, referral string) (resp response.DataHome, err error)

	//product
	AddProductSales(ctx context.Context, param request.AddProductReq) (err error)
	ListProductSales(ctx context.Context, paramFilter sharedModel.QueryRequest) (resp []response.ProductListBySales, pagination *sharedModel.Pagination, err error)
	GetDetailSalesProduct(ctx context.Context, id string) (resp response.ProductDetailResp, err error)
	UpdateProductSales(ctx context.Context, param request.UpdateProductSales) (err error)

	//banner
	AddBannerSales(ctx context.Context, param request.BannerReq) (err error)
	GetListBannerSales(ctx context.Context) (resp response.BannerListSalesResp, err error)
	GetListBannerPublic(ctx context.Context, subdomain, referral string) (resp response.BannerListPublicSalesResp, err error)
	GetBannerSales(ctx context.Context, id string) (resp response.BannerResp, err error)
	UpdateBanner(ctx context.Context, req request.BannerUpdateReq) (err error)

	//gallery
	AddGallerySales(ctx context.Context, param request.AddGalleryParam) (err error)
	GetDataListGallery(ctx context.Context) (resp response.GalleryResp, err error)
	GetDataListGalleryPublic(ctx context.Context, subdomain, referral string) (resp response.GalleryPublicResp, err error)
	GetDataGallerySales(ctx context.Context, id string) (resp response.GalleryDataResp, err error)
	UpdateGallery(ctx context.Context, req request.UpdateGalleryParam) (err error)

	//social_media
	AddSocialMediaSales(ctx context.Context, req request.AddSocialMediaReq) (err error)
	GetListSocialMediaSales(ctx context.Context) (resp response.SocialMediaListResp, err error)
	GetListSocialMediaPublicSales(ctx context.Context, subdomain, referral string) (resp response.SocialMediaListResp, err error)

	//template
	AddTemplateSales(ctx context.Context, param request.AddTemplateReq) (err error)
	GetListTemplateSales(ctx context.Context) (resp response.TemplateListSalesResp, err error)
	GetListPublicTemplateSales(ctx context.Context, subdomain, referral string) (resp response.TemplateListPublicResp, err error)
	GetDetailTemplateSales(ctx context.Context, id string) (resp response.TemplateDetailResp, err error)
	UpdateTemplateSales(ctx context.Context, req request.UpdateTemplateReq) (err error)

	//testimony
	AddTestimoniFeature(ctx context.Context, request request.Testimoni) (err error)
	UpdateTestimoniFeature(ctx context.Context, request request.Testimoni) (err error)
	GetTestimoniFeature(ctx context.Context, request request.Testimoni) (response request.Testimoni, err error)
	GetListTestimoniFeature(ctx context.Context, request request.Testimoni, filter request.TestimoniSearch) (response []request.Testimoni, pagination *sharedModel.Pagination, err error)
}

type salesFeature struct {
	repo     repository.SalesRepository
	userRepo repoUser.UserRepository
}

func NewSalesFeature(repo repository.SalesRepository, userRepo repoUser.UserRepository) SalesFeature {
	return &salesFeature{
		repo:     repo,
		userRepo: userRepo,
	}
}
