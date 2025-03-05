package http

import (
	"fmt"
	"zayyid-go/delivery/http/middleware"
	"zayyid-go/infrastructure/helper"

	"github.com/gofiber/fiber/v2"
)

// ===============================================================================
// Only single route defined to handle all request from third party apps (source)
// Any http method would be allowed but will be validated on the handler
// `:base_path` is app id param, `*` is param to capture any endpoint
// ===============================================================================
// Example https://app.company.com/middleware/wms-dc-app/v1/plant
// BASEURL https://app.company.com/middleware
// APPID wmd-dc-app
// ENDPOINT /v1/plant
// ===============================================================================
func RegisterRoute(app *fiber.App, handler Handler) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		hwHealth := helper.NewHwStats()
		return c.Status(200).JSON(fiber.Map{"cpu_cores": hwHealth.NumCpu, "cpu_usage": fmt.Sprintf("%.2f%%", hwHealth.CpuUsage), "mem_total": hwHealth.TotalMem, "mem_avail": hwHealth.AvailMem, "db_status": "health", "tgt_status": "health"})
	})

	// app.Get("/signature", handler.generateSignatureHandler.Handle)
	// app.Post("/signature", handler.generateSignatureHandler.Handle)

	api := app.Group("/api")
	{
		api.Get("/ping", handler.userMenuHandler.Ping)
	}

	thirdParty := app.Group("/third-party")
	{
		thirdParty.Post("/callback-payment-receiving", handler.thirdPartyHandler.CallbackPaymentReceivingMidtrans)
		thirdParty.Post("/frontend-payment-notification", handler.thirdPartyHandler.FrontendPaymentNotification)
	}

	public := app.Group("/public")
	{
		public.Get("/home/:subdomain/*", handler.salesHandler.GetDataHome)
		public.Get("/gallery/:subdomain", handler.salesHandler.GetGallerySalesPublic)
		public.Get("/banner/:subdomain", handler.salesHandler.GetBannerPublicSales)
		public.Get("/template/:subdomain", handler.salesHandler.GetListPublicTemplateSales)
		public.Get("/social-media/:subdomain", handler.salesHandler.GetListPublicSocialMediaSales)
		public.Get("/testimony/:subdomain/*", handler.salesHandler.GetPublicListTestimoniHandler)
		public.Get("/product/:subdomain", handler.salesHandler.GetListProductSalesPublic)
		public.Get("/product/detail/:subdomain/:slug", handler.salesHandler.GetDetailProductSalesPublic)
	}

	master := app.Group("/master")
	{
		master.Get("/ping", handler.masterHandler.PingMaster)
		master.Get("/province", handler.masterHandler.MasterListProvince)
		master.Get("/city", handler.masterHandler.MasterListCity)
	}

	sales := app.Group("/sales")
	{
		sales.Post("/product", middleware.Auth, handler.salesHandler.AddProductSales)
		sales.Get("/product", middleware.Auth, handler.salesHandler.GetListProductSales)
		sales.Get("/product/:id", middleware.Auth, handler.salesHandler.GetDetailProductSales)
		sales.Put("/product/:id", middleware.Auth, handler.salesHandler.UpdateProductSales)

		//gallery
		sales.Post("/gallery", middleware.Auth, handler.salesHandler.AddGallerySales)
		sales.Get("/gallery", middleware.Auth, handler.salesHandler.GetGallerySales)
		sales.Get("/gallery/:id", middleware.Auth, handler.salesHandler.GetDataGallerySales)
		sales.Put("/gallery/:id", middleware.Auth, handler.salesHandler.UpdateGallerySales)

		//testimony
		sales.Get("/testimony", handler.salesHandler.GetTestimoniHandler)
		sales.Get("/testimony/list", handler.salesHandler.GetListTestimoniHandler)
		sales.Post("/testimony", handler.salesHandler.AddTestimoniHandler)
		sales.Put("/testimony", handler.salesHandler.UpdateTestimoniHandler)

		//banner
		sales.Post("/banner", middleware.Auth, handler.salesHandler.AddBannerSales)
		sales.Get("/banner", middleware.Auth, handler.salesHandler.GetListBannerSales)
		sales.Get("/banner/:id", middleware.Auth, handler.salesHandler.GetBannerSales)
		sales.Put("/banner/:id", middleware.Auth, handler.salesHandler.UpdateBannerSales)

		//template
		sales.Post("/template", middleware.Auth, handler.salesHandler.AddTemplateSales)
		sales.Get("/template", middleware.Auth, handler.salesHandler.GetListTemplateSales)
		sales.Get("/template/:id", middleware.Auth, handler.salesHandler.GetTemplateSales)
		sales.Put("/template/:id", middleware.Auth, handler.salesHandler.UpdateTemplateSales)

		//social media
		sales.Post("/social-media", middleware.Auth, handler.salesHandler.AddSocialMediaSales)
		sales.Get("/social-media", middleware.Auth, handler.salesHandler.GetListSocialMediaSales)
		sales.Get("/social-media/:id", middleware.Auth, handler.salesHandler.GetSocialMediaSales)
		sales.Put("/social-media/:id", middleware.Auth, handler.salesHandler.UpdateSocialMediaSales)
	}

	agent := app.Group("/agent")
	{
		agent.Post("/create", middleware.Auth, handler.userHandler.CreateAgentHandler)
		agent.Get("/list", middleware.Auth, handler.userHandler.CreateAgentHandler)
	}

	// user endpoint
	user := app.Group("/user")
	{
		user.Post("/register", handler.userHandler.RegisterUserHandler)
		user.Post("/login", handler.userHandler.AuthUserHandler)
		user.Post("/refresh-token", handler.userHandler.RefreshTokenHandler)
		user.Put("/update", middleware.Auth, handler.userHandler.UpdateHandler)
	}

	// app.All("/:base_path/*", middleware.SignatureMiddleware(), handler.applicationMenuHandler.Handle)
}
