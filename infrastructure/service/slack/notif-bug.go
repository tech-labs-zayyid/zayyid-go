package slack

import (
	libSlack "github.com/slack-go/slack"
)

type ConfigSlack struct {
	ApiToken  string
	ChannelId string
}

type SlackNotificationBug interface {
	Send(message string) error
}

type slack struct {
	config ConfigSlack
}

func InitConnectionSlack(conf ConfigSlack) SlackNotificationBug {
	return &slack{
		config: conf,
	}
}

func (s slack) Send(message string) error {
	api := libSlack.New(s.config.ApiToken)
	_, _, err := api.PostMessage(s.config.ChannelId, libSlack.MsgOptionText(message, false))
	return err
}
