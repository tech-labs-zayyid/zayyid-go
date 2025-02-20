package config

import (
	"fmt"
	"os"
	"strconv"

	"zayyid-go/infrastructure/database"
	"zayyid-go/infrastructure/service/slack"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
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
		Slack: slack.ConfigSlack{
			ApiToken:  os.Getenv("API_TOKEN"),
			ChannelId: os.Getenv("CHANNEL_ID"),
		},
	}

	return
}
