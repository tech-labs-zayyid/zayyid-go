package handler

import (
	"zayyid-go/domain/sales/feature"
	"zayyid-go/infrastructure/service/slack"

	"github.com/gofiber/fiber/v2"
)

type salesHandler struct {
	feature         feature.SalesFeature
	isRequestLogged bool
	slackConf       slack.SlackNotificationBug
}

func NewSalesHandler(feature feature.SalesFeature, isRequestLogged bool, slackConfig slack.SlackNotificationBug) SalesHandlerInterface {
	return salesHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
		slackConf:       slackConfig,
	}
}

type SalesHandlerInterface interface {
	GetDataHome(c *fiber.Ctx) (err error)

	//product
	AddProductSales(c *fiber.Ctx) (err error)

	//gallery
	AddGallerySales(c *fiber.Ctx) (err error)
	GetGallerySales(c *fiber.Ctx) (err error)
	GetGallerySalesPublic(c *fiber.Ctx) (err error)
	GetDataGallerySales(c *fiber.Ctx) (err error)
	UpdateGallerySales(c *fiber.Ctx) (err error)

	// testimony
	AddTestimoniHandler(c *fiber.Ctx) error
	UpdateTestimoniHandler(c *fiber.Ctx) error
	GetTestimoniHandler(c *fiber.Ctx) error
	GetListTestimoniHandler(c *fiber.Ctx) error
	GetPublicListTestimoniHandler(c *fiber.Ctx) error

	// banner
	AddBannerSales(c *fiber.Ctx) (err error)
	GetListBannerSales(c *fiber.Ctx) (err error)
	GetBannerPublicSales(c *fiber.Ctx) (err error)
	GetBannerSales(c *fiber.Ctx) (err error)
	UpdateBannerSales(c *fiber.Ctx) (err error)

	//template
	AddTemplateSales(c *fiber.Ctx) (err error)
	GetListTemplateSales(c *fiber.Ctx) (err error)
	GetListPublicTemplateSales(c *fiber.Ctx) (err error)
	GetTemplateSales(c *fiber.Ctx) (err error)
	UpdateTemplateSales(c *fiber.Ctx) (err error)

	//social media
	AddSocialMediaSales(c *fiber.Ctx) (err error)
	GetListSocialMediaSales(c *fiber.Ctx) (err error)
	GetListPublicSocialMediaSales(c *fiber.Ctx) (err error)
}
