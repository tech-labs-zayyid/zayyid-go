package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"middleware-cms-api/infrastructure/logger"
	sharedConstant "middleware-cms-api/infrastructure/shared/constant"
	logError "middleware-cms-api/infrastructure/shared/error"
	"time"

	"github.com/streadway/amqp"
)

func (q queueService) PublishData(ctx context.Context, topic string, msg interface{}) (err error) {

	cfg := q.rabbitmq.GetConfig()

	select {
	case err := <-cfg.Err:
		if err != nil {
			q.rabbitmq.Reconnect()
		}
	default:
	}

	ch, err := cfg.Conn.Channel()
	if err != nil {
		err = logError.New(sharedConstant.PUBLISHER_RABBITMQ, sharedConstant.ERR_DEFINE_CHANNEL_TO_BROKER, err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		topic,                                 // name
		sharedConstant.RABBITMQ_EXCHANGE_TYPE, // type
		true,                                  // durable
		false,                                 // auto-deleted
		false,                                 // internal
		false,                                 // no-wait
		nil,                                   // arguments
	)
	if err != nil {
		err = logError.New(sharedConstant.PUBLISHER_RABBITMQ, sharedConstant.ERR_DECLARE_EXCHANGE_TO_BROKER, err)
		return
	}

	body, err := json.Marshal(msg)
	if err != nil {
		err = logError.New(sharedConstant.PUBLISHER_RABBITMQ, sharedConstant.ERR_MARSHAL_JSON, err)
		return
	}

	err = cfg.Channel.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,    // keeping message if broker restart
			ContentType:  "application/json", // XXX: We will revisit this in future episodes
			Body:         body,
			Timestamp:    time.Now(),
		})
	if err != nil {
		err = logError.New(sharedConstant.PUBLISHER_RABBITMQ, sharedConstant.ERR_PUBLISH_QUEUE_TO_BROKER, err)
		return
	}

	fmt.Println(fmt.Sprintf(sharedConstant.SUCCESS_PUBLISH_TO_BROKER, topic))
	logger.LogInfo(sharedConstant.PUBLISHER_RABBITMQ, fmt.Sprintf(sharedConstant.SUCCESS_PUBLISH_TO_BROKER, topic))

	return

}
