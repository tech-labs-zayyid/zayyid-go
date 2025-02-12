package container

import (
	"log"

	"middleware-cms-api/config"
	"middleware-cms-api/delivery/cron"
	atomicRepo "middleware-cms-api/domain/shared/repository"
	UserMenu "middleware-cms-api/domain/user_menu/feature"
	UserRepo "middleware-cms-api/domain/user_menu/repository"
	"middleware-cms-api/infrastructure/database"
	"middleware-cms-api/infrastructure/logger"
	"middleware-cms-api/infrastructure/minio"
	"middleware-cms-api/infrastructure/service/queue"
)

type Container struct {
	Database          *database.Database
	QueueServices     queue.QueueService
	EnvironmentConfig config.EnvironmentConfig
	UserMenuFeature   *UserMenu.UserMenuFeature
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
		),
	}
}
