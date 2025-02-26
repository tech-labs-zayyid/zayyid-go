package container

import (
	"log"

	"zayyid-go/config"
	Master "zayyid-go/domain/master/feature"
	MasterRepo "zayyid-go/domain/master/repository"
	Sales "zayyid-go/domain/sales/feature"
	SalesRepo "zayyid-go/domain/sales/repository"
	atomicRepo "zayyid-go/domain/shared/repository"
	User "zayyid-go/domain/user/feature"
	UserRepo "zayyid-go/domain/user/repository"
	UserMenu "zayyid-go/domain/user_menu/feature"
	UserMenuRepo "zayyid-go/domain/user_menu/repository"
	"zayyid-go/infrastructure/database"
	"zayyid-go/infrastructure/logger"
	"zayyid-go/infrastructure/service/queue"
	"zayyid-go/infrastructure/service/slack"
)

type Container struct {
	Database          *database.Database
	QueueServices     queue.QueueService
	EnvironmentConfig config.EnvironmentConfig
	UserMenuFeature   *UserMenu.UserMenuFeature
	MasterFeature     *Master.MasterFeature
	SalesFeature      *Sales.SalesFeature
	Slack             *slack.ConfigSlack
	UserFeature User.UserFeature
}

func SetupContainer() Container {
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	logger.InitializeLogger(logger.LOGRUS) // choose which log, ZAP or LOGRUS. Default: LOGRUS

	db, err := database.LoadDatabase(config.Database)
	if err != nil {
		log.Panic(err)
	}

	//cron.Run()

	notifBug := slack.InitConnectionSlack(config.Slack)

	return Container{
		EnvironmentConfig: config,
		UserMenuFeature: UserMenu.NewUserMenuFeature(
			config,
			UserMenuRepo.NewUserMenuRepository(db),
			atomicRepo.NewUOWRepository(db),
			notifBug,
		),
		MasterFeature: Master.NewMasterFeature(
			config,
			MasterRepo.NewMasterRepository(db),
			atomicRepo.NewUOWRepository(db),
			notifBug,
		),
		SalesFeature: Sales.NewSalesFeature(
			config,
			SalesRepo.NewSalesRepository(db),
			notifBug,
		),
		UserFeature: User.NewUserFeature(
			UserRepo.NewUserRepository(db),
		),
	}
}
