package queue

import (
	"context"
	"log"

	"fmt"
	"middleware-cms-api/infrastructure/logger"
	sharedConstant "middleware-cms-api/infrastructure/shared/constant"
	logError "middleware-cms-api/infrastructure/shared/error"

	"github.com/streadway/amqp"
)

func checkTopic(routeKey string, exchangeKey string, topicValue string) (response bool) {
	log.Println("routeKey : ", routeKey)
	log.Println("exchangeKey : ", exchangeKey)
	log.Println("topicValue : ", topicValue)
	if routeKey != topicValue {
		if exchangeKey == topicValue {
			response = true
		}
	} else {
		if routeKey == topicValue || exchangeKey == topicValue {
			response = true
		}
	}

	return
}

func (q queueService) ConsumeData(ctx context.Context, topic string) (err error) {
	cfg := q.rabbitmq.GetConfig()
	notify := cfg.Conn.NotifyClose(make(chan *amqp.Error)) //error channel

	ch, err := cfg.Conn.Channel()
	if err != nil {
		err = logError.New("consumer", sharedConstant.ERR_CREATE_CHANNEL_TO_BROKER, err)
		return
	}

	// NOTE : kalo di local uncomment, kalo mau push comment dulu
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
		return
	}

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		err = logError.New(sharedConstant.CONSUMER_BILLING_RABBITMQ, sharedConstant.ERR_CREATE_QUEUE_TO_BROKER, err)
		return
	}

	err = cfg.Channel.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		topic,      // exchange
		false,
		nil,
	)
	if err != nil {
		err = logError.New(sharedConstant.CONSUMER_BILLING_RABBITMQ, sharedConstant.ERR_BINDING_QUEUE_TO_BROKER, err)
		return
	}

	// consume
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack ini untuk falging apakah data telah diterima atau belum
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		err = logError.New(sharedConstant.CONSUMER_BILLING_RABBITMQ, sharedConstant.ERR_CONSUME_QUEUE_TO_BROKER, err)
		return
	}

	logger.LogInfo(sharedConstant.CONSUMER_BILLING_RABBITMQ, fmt.Sprintf(sharedConstant.START_LISTENING_TOPIC_FROM_BROKER, topic))

	for {
		select {
		case err = <-notify:
			// if err != nil {
			errN := err // handle race condition
			log.Println(logError.New(sharedConstant.CONSUMER_BILLING_RABBITMQ+" errN", sharedConstant.ERR_CONSUME_QUEUE_TO_BROKER, errN))
			for {
				err = q.rabbitmq.Reconnect()
				if err == nil {
					break
				}
			}
			// }
		case msg := <-msgs:
			if checkTopic(msg.RoutingKey, msg.Exchange, "producer_name") {
				logger.LogInfo(sharedConstant.CONSUMER_BILLING_RABBITMQ, fmt.Sprintf(sharedConstant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))
				// do something
				q.PublishData(ctx, "producer_name", []interface{}{})
			} else if checkTopic(msg.RoutingKey, msg.Exchange, "consumer_name") {
				logger.LogInfo(sharedConstant.CONSUMER_BILLING_RABBITMQ, fmt.Sprintf(sharedConstant.SUCCESS_CONSUME_FROM_BROKER, topic, string(msg.Body)))
				// do something
				q.PublishData(ctx, "consumer_name", []interface{}{})
			}
		}
	}
}
