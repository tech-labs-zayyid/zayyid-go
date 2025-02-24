package http

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"zayyid-go/delivery/container"
	"zayyid-go/delivery/http/middleware"
	errHelper "zayyid-go/domain/shared/helper/error"
	sharedResponse "zayyid-go/domain/shared/response"
	"zayyid-go/infrastructure/logger"

	_ "zayyid-go/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func ServeHttp(container container.Container) *fiber.App {
	handler := SetupHandler(container)
	isPreforkEnabled, err := strconv.ParseBool(os.Getenv("ENABLE_APP_PREFORK"))
	if err != nil {
		fmt.Println("App prefork is disabled due to misconfiguration in ENABLE_APP_PREFORK environment variable")
	}

	app := fiber.New(
		fiber.Config{
			Prefork:      isPreforkEnabled,
			ErrorHandler: handleGlobalError,
		},
	)

	config := recover.Config{
		EnableStackTrace: true,
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Use(recover.New(config), cors.New(), middleware.Timer())
	app.Get("/swagger/*", swagger.HandlerDefault)
	RegisterRoute(app, handler)

	return app
}

// Catch all errors and return as proper response
func handleGlobalError(c *fiber.Ctx, err error) error {
	requestUrl := c.Request().URI().String()
	log.Println("requestUrl : ", requestUrl)
	log.Println("err : ", err.Error())
	a, b := err.(errHelper.IntegrationError)
	log.Println("err.(errHelper.IntegrationError) : ", a)
	log.Println("err.(errHelper.IntegrationError) b : ", b)
	if iErr, ok := err.(errHelper.IntegrationError); !ok {
		log.Println("iErr : ", iErr.Message)
		log.Println("iErr type : ", iErr.Type)
		log.Println("ok : ", ok)
		switch iErr.Type {
		case errHelper.ERROR_NOT_AUTHORIZED:
			return sharedResponse.NotAuthorizedError(c, iErr.Message, requestUrl)
		case errHelper.ERROR_DATA_NOT_FOUND:
			return sharedResponse.NotFoundError(c, iErr.Message, requestUrl, nil)
		case errHelper.ERROR_NETWORK_CONNECTION:
			return sharedResponse.NetworkConnectionError(c, iErr.Message, requestUrl)
		case errHelper.ERROR_REQUEST_TIMEOUT:
			return sharedResponse.RequestTimeoutError(c, iErr.Message, requestUrl)
		default:
			// All not catcable error would be considered as internal server error below
		}
	}

	log.Println("lewat")
	// Non integration error would be treated as internal server error
	logger.LogError("unhandled_error", "internal_server_error", err.Error())
	return sharedResponse.InternalServerError(c, "Internal server error", requestUrl, nil)
}
