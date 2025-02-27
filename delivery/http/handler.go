package http

import (
	"os"
	"strconv"
	"zayyid-go/delivery/container"
	MasterHandler "zayyid-go/domain/master/handler"
	SalesHandler "zayyid-go/domain/sales/handler"
	ThirdPartyHandler "zayyid-go/domain/third_party/handler"
	UserHandler "zayyid-go/domain/user/handler"
	UserMenuHandler "zayyid-go/domain/user_menu/handler"
)

type Handler struct {
	userMenuHandler   UserMenuHandler.UserHandlerInterface
	masterHandler     MasterHandler.MasterHandlerInterface
	salesHandler      SalesHandler.SalesHandlerInterface
	userHandler       UserHandler.IUserHandler
	thirdPartyHandler ThirdPartyHandler.ThirdPartyHandler
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
		masterHandler: MasterHandler.NewMasterHandler(
			container.MasterFeature, isRequestLogged,
		),
		salesHandler: SalesHandler.NewSalesHandler(
			container.SalesFeature, isRequestLogged,
		),
		userHandler: UserHandler.NewUserHandler(
			&container.UserFeature,
		),
		thirdPartyHandler: ThirdPartyHandler.NewThirdPartyHandler(
			&container.ThirdPartyFeature,
		),
	}
}
