package rabbitmq

import (
	"fmt"
	sharedConstant "zayyid-go/infrastructure/shared/constant"

	"github.com/streadway/amqp"
)

type RabbitmqConfig struct {
	Host         string
	Username     string
	Password     string
	Port         int
	ProducerName string
	ConsumerName string
}

type RabbitMQ interface {
	Connect() (err error)
	Close()
	Reconnect() error
	GetConfig() rabbitMQ
}

type rabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Err     chan error
	config  RabbitmqConfig
}

func NewConnection(config RabbitmqConfig) RabbitMQ {
	return &rabbitMQ{
		config: config,
		Err:    make(chan error),
	}
}

func (c *rabbitMQ) GetConfig() rabbitMQ {
	return *c
}

func (c *rabbitMQ) Connect() (err error) {
	connPattern := "amqp://%v:%v@%v:%v"
	if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		c.config.Username,
		c.config.Password,
		c.config.Host,
		c.config.Port,
	)

	if c.config.Port == 0 {
		connPattern = "amqp://%v:%v@%v"
		clientUrl = fmt.Sprintf(connPattern,
			c.config.Username,
			c.config.Password,
			c.config.Host,
		)
	} else if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	c.Conn, err = amqp.Dial(clientUrl)
	if err != nil {
		fmt.Println(err)
		if err = c.Retry(); err != nil {
			err = fmt.Errorf(sharedConstant.ERR_CONN_TO_BROKER, err)
		}
	}

	c.Channel, err = c.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf(sharedConstant.ERR_CREATE_CHANNEL_TO_BROKER, err)
		return
	}

	if err = c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		err = fmt.Errorf(sharedConstant.ERR_SETUP_QUEUE_TO_BROKER, err)
		return
	}

	return
}

func (c *rabbitMQ) Retry() (err error) {
	connPattern := "amqp://%v:%v@%v:%v"
	if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		c.config.Username,
		c.config.Password,
		c.config.Host,
		c.config.Port,
	)

	if c.config.Port == 0 {
		connPattern = "amqp://%v:%v@%v"
		clientUrl = fmt.Sprintf(connPattern,
			c.config.Username,
			c.config.Password,
			c.config.Host,
		)
	} else if c.config.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	conn, err := amqp.Dial(clientUrl)
	if err != nil {
		err = fmt.Errorf(sharedConstant.ERR_CONN_TO_BROKER, err)
		return
	}

	c.Conn = conn

	return
}

func (c *rabbitMQ) Close() {
	c.Conn.Close()
}

func (c *rabbitMQ) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}
