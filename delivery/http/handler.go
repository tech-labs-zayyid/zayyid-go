package http

import (
	"os"
	"strconv"
	"zayyid-go/delivery/container"
	MasterHandler "zayyid-go/domain/master/handler"
	SalesHandler "zayyid-go/domain/sales/handler"
	UserHandler "zayyid-go/domain/user/handler"
	UserMenuHandler "zayyid-go/domain/user_menu/handler"
	"zayyid-go/infrastructure/service/slack"
)

type Handler struct {
	slackConfig     slack.SlackNotificationBug
	userMenuHandler UserMenuHandler.UserHandlerInterface
	masterHandler   MasterHandler.MasterHandlerInterface
	salesHandler    SalesHandler.SalesHandlerInterface
	userHandler     UserHandler.IUserHandler
}

func SetupHandler(container container.Container) Handler {
	isRequestLogged, err := strconv.ParseBool(os.Getenv("ENABLE_REQUEST_LOG"))
	if err != nil {
		isRequestLogged = false
	}

	notifBug := slack.InitConnectionSlack(container.EnvironmentConfig.Slack)
	return Handler{
		userMenuHandler: UserMenuHandler.NewUserMenuHandler(
			container.UserMenuFeature, isRequestLogged,
		),
		masterHandler: MasterHandler.NewMasterHandler(
			container.MasterFeature, isRequestLogged,
		),
		salesHandler: SalesHandler.NewSalesHandler(
			container.SalesFeature, isRequestLogged, notifBug,
		),
		userHandler: UserHandler.NewUserHandler(
			&container.UserFeature,
		),
	}
}
