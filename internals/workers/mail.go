package workers

import (
	"context"
	"fmt"

	"github.com/manasss0508/notifyx-go/internals/api"
	"github.com/manasss0508/notifyx-go/internals/repository"
	"github.com/manasss0508/notifyx-go/internals/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

// takes consumer and process each message
func ProcessEachMessageMail(state *api.AppState, consumer <-chan amqp.Delivery) {
	for msg := range consumer {
		//fmt.Println(msg)
		// starting go routine to process message
		go func() {
			// getting state
			state := state

			// getting message
			message := msg

			// deserialize message to struct and get notification
			notification, err := service.DeserializeMessage(message.Body)
			if err != nil {
				fmt.Println(err)
			}
			notificationId := notification.NotficationId

			// get notification from database
			notif, err := repository.DbGetNotification(context.Background(), (*state).OtelTracer, (*state).Db, notificationId)
			if err != nil {
				fmt.Println(err)
				return
			}

			// get template
			template, err := (*(*state).TemplateCache).GetTemplateMail((*state).Db, notif.Channel, notif.Template)
			if err != nil {
				fmt.Println(err)
				return
			}

			// render template
			subject, body := service.TemplateRender(template, notif.Variables)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(subject, body)

			// send main

			// acknowledge message
			message.Acknowledger.Ack(0, false)
		}()
	}
}
