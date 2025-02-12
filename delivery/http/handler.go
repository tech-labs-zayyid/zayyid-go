package http

import (
	"middleware-cms-api/delivery/container"
	UserMenuHandler "middleware-cms-api/domain/user_menu/handler"
	"os"
	"strconv"
)

type Handler struct {
	userMenuHandler UserMenuHandler.UserHandlerInterface
}

func SetupHandler(container container.Container) Handler {
	isRequestLogged, err := strconv.ParseBool(os.Getenv("ENABLE_REQUEST_LOG"))
	if err != nil {
		isRequestLogged = false
	}

	return Handler{
		userMenuHandler: UserMenuHandler.NewUserMenuHandler(
			container.UserMenuFeature, isRequestLogged,
		),
	}
}
