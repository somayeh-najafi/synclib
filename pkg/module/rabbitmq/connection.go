package rabbitmq

import (
	"github.com/qmstr/synclib/module/util"
	"github.com/streadway/amqp"
)

// Connects to a RabbitMQ instance
func Connect(rabbitMqAddress string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(rabbitMqAddress)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	return conn, ch
}
