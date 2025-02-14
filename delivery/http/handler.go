package http

import (
	"os"
	"strconv"
	"zayyid-go/delivery/container"
	UserMenuHandler "zayyid-go/domain/user_menu/handler"
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
