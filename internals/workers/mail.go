package workers

import (
	"fmt"

	"github.com/manasss0508/notifyx-go/internals/api"
	amqp "github.com/rabbitmq/amqp091-go"
)

// takes consumer and process each message
func ProcessEachMessageMail(state api.AppState, consumer <-chan amqp.Delivery) {
	for msg := range consumer {
		fmt.Println(msg)
	}
}
