package container

import (
	"log"

	"zayyid-go/config"
	"zayyid-go/delivery/cron"
	atomicRepo "zayyid-go/domain/shared/repository"
	UserMenu "zayyid-go/domain/user_menu/feature"
	UserRepo "zayyid-go/domain/user_menu/repository"
	"zayyid-go/infrastructure/database"
	"zayyid-go/infrastructure/logger"
	"zayyid-go/infrastructure/minio"
	"zayyid-go/infrastructure/service/queue"
	"zayyid-go/infrastructure/service/slack"
)

type Container struct {
	Database          *database.Database
	QueueServices     queue.QueueService
	EnvironmentConfig config.EnvironmentConfig
	UserMenuFeature   *UserMenu.UserMenuFeature
	Slack             *slack.ConfigSlack
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

	cron.Run()

	_, err = minio.MinioConnection(config)
	if err != nil {
		log.Panic(err)
	}

	notifBug := slack.InitConnectionSlack(config.Slack)

	// rmq := rabbitmq.NewConnection(config.RabbitMq)
	// // Connect RabbitMQ
	// err = rmq.Connect()
	// if err != nil {
	// 	log.Panic(err)
	// }

	return Container{
		EnvironmentConfig: config,
		UserMenuFeature: UserMenu.NewUserMenuFeature(
			config,
			UserRepo.NewUserMenuRepository(db),
			atomicRepo.NewUOWRepository(db),
			notifBug,
		),
	}
}
