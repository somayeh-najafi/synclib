package rabbitmq

import (
	"fmt"
	"github.com/qmstr/synclib/module/util"
	"github.com/streadway/amqp"
	"log"
)

type Callback func()

func OnMessageReceive(ch *amqp.Channel, queueName string, callback Callback) chan bool {

	// Consuming the order message
	msgs, err := ch.Consume(
		queueName, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	// Sending back the module result
	responded := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Executing the module's callback function
			log.Printf("Executing callback function...")
			callback()

			// Sending a response
			log.Printf("...callback function completed. Sending response...")
			responseBody := fmt.Sprintf("Response from %v", queueName)
			err = ch.Publish(
				"",        			// exchange
				ResponseQueueName(queueName),	// routing key
				false,     			// mandatory
				false,     			// immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: queueName,
					Body:          []byte(responseBody),
				})
			util.FailOnError(err, "Failed to publish the response")
			log.Printf("...response sent.")
			responded <- true
		}
	}()

	// Returning Go channel on which the caller will have to wait
	return responded
}
