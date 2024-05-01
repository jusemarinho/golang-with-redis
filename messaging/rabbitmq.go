package messaging

import (
	"fmt"
	"golang-with-redis/config"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbitMQ(amc *config.AmqpConfig) (*RabbitMQ, error) {

	strConnection := fmt.Sprintf("%s://%s:%s@%s:%d", amc.Protocol, amc.User, amc.Password, amc.Host, amc.Port)

	conn, err := amqp.Dial(strConnection)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"queue_pizzas",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		connection: conn,
		channel:    ch,
		queue:      q,
	}, nil
}

func (rmq *RabbitMQ) PublishMessage(body string) error {
	err := rmq.channel.Publish(
		"",
		rmq.queue.Name,
		false,
		false,
		amqp.Publishing{
			Body: []byte(body),
			Headers: amqp.Table{
				"header-test": "application/json",
			},
		})
	if err != nil {
		return err
	}
	return nil
}
