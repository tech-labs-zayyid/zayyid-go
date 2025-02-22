package http

import (
	"fmt"

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

	public := app.Group("public")
	{
		public.Get("/gallery/:subdomain", handler.salesHandler.GetGallerySalesPublic)
	}

	master := app.Group("/master")
	{
		master.Get("/ping", handler.masterHandler.PingMaster)
		master.Get("/province", handler.masterHandler.MasterListProvince)
		master.Get("/city", handler.masterHandler.MasterListCity)
	}

	sales := app.Group("/sales")
	{
		sales.Post("/gallery", handler.salesHandler.AddGallerySales)
		sales.Get("/gallery", handler.salesHandler.GetGallerySales)
		sales.Get("/testimony", handler.salesHandler.GetTestimoniHandler)
		sales.Get("/testimony/list", handler.salesHandler.GetListTestimoniHandler)
		sales.Post("/testimony", handler.salesHandler.AddTestimoniHandler)
		sales.Put("/testimony", handler.salesHandler.UpdateTestimoniHandler)
	}

	_ = app.Group("/agent")
	{
		//list API for Agent
	}

	// app.All("/:base_path/*", middleware.SignatureMiddleware(), handler.applicationMenuHandler.Handle)
}
