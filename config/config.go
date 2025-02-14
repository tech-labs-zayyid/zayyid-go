package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"zayyid-go/infrastructure/broker/rabbitmq"
	"zayyid-go/infrastructure/database"
	"zayyid-go/infrastructure/service/slack"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig

	StorageMinioServer     string
	StorageMinioAccessKey  string
	StorageMinioSecreatKey string
	StorageMinioBucket     string
	StorageMinioUseSSL     string
	StorageMinioPMAServer  string
	StorageMinioPMAGateway string

	RabbitMq rabbitmq.RabbitmqConfig
	Slack    slack.ConfigSlack
}

type App struct {
	Name    string
	Version string
	Port    int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf("failed to load env: %s", err.Error())
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf("invalid APP_PORT config: %s", err.Error())
		return
	}

	rmqPort := 0
	if os.Getenv("RABBITMQ_PORT") != "" {
		rmqPort, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
		if err != nil {
			err = fmt.Errorf("error convert string to int", err)
			return
		}
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.DatabaseConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		StorageMinioServer:     os.Getenv("STORAGE_MINIO_SERVER"),
		StorageMinioAccessKey:  os.Getenv("STORAGE_MINIO_ACCESS_KEY"),
		StorageMinioSecreatKey: os.Getenv("STORAGE_MINIO_SECRET_KEY"),
		StorageMinioBucket:     os.Getenv("STORAGE_MINIO_BUCKET"),
		StorageMinioUseSSL:     os.Getenv("STORAGE_MINIO_USE_SSL"),
		StorageMinioPMAServer:  os.Getenv("STORAGE_MINIO_PMA_SERVER"),
		StorageMinioPMAGateway: os.Getenv("STORAGE_MINIO_PMA_GATEWAY"),
		RabbitMq: rabbitmq.RabbitmqConfig{
			Host:         os.Getenv("RABBITMQ_HOST"),
			Username:     os.Getenv("RABBITMQ_USERNAME"),
			Password:     os.Getenv("RABBITMQ_PASSWORD"),
			Port:         rmqPort,
			ProducerName: os.Getenv("RABBITMQ_PRODUCER_NAME"),
			ConsumerName: os.Getenv("RABBITMQ_CONSUMER_NAME"),
		},
		Slack: slack.ConfigSlack{
			ApiToken:  os.Getenv("API_TOKEN"),
			ChannelId: os.Getenv("CHANNEL_ID"),
		},
	}

	return
}
