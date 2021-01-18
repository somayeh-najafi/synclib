package rabbitmq

import (
	"fmt"
	"github.com/qmstr/synclib/module/util"
	"github.com/streadway/amqp"
	"log"
)

// Declares two queues:
// - "queueName" contains orders from the Orchestrator to the module
// - "queueName-response" contains response messages from the module to the Orchestrator
func DeclareQueues(ch *amqp.Channel, queueName string) (amqp.Queue, amqp.Queue) {
	return declareQueue(ch, queueName), declareResponseQueue(ch, queueName)
}

// Given a queue name, returns its response queue name
func ResponseQueueName(queueName string) string {
	return fmt.Sprintf("%v-response", queueName)
}

// Declares a single "response" RabbitMQ queue,
// i.e., a queue with the "-response" suffix
func declareResponseQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	return declareQueue(ch, ResponseQueueName(queueName))
}

// Declares a single RabbitMQ queue
func declareQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	log.Printf("Declaring queue \"%s\"", queueName)
	q, err := ch.QueueDeclare(
		queueName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	util.FailOnError(err, "Failed to declare queue \"%s\", queueName")

	return q
}
