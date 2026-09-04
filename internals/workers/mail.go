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
		fmt.Println("notification recived : ", msg)
		// starting go routine to process message
		go processMessage(state, msg)
	}
	fmt.Println("🚨 CONSUMER CHANNEL CLOSED")
}

func processMessage(state *api.AppState, message amqp.Delivery) {
	// deserialize message to struct and get notification
	notification, err := service.DeserializeMessage(message.Body)
	if err != nil {
		fmt.Println(err)
		return
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

	// send mail
	err = (*state).EmailService.Send(notif.Recipient, subject, body)

	// nack if mail not sent
	if err != nil {
		fmt.Println("mail failed to send : ", err)

		// nack message
		nackErr := message.Acknowledger.Nack(
			message.DeliveryTag,
			false,
			false,
		)
		if nackErr != nil {
			fmt.Println("NACK failed:", nackErr)
		}

		// update status in db
		repository.DbUpdateNotificationStatus(
			context.Background(),
			(*state).OtelTracer,
			(*state).Db,
			notificationId,
			"FAILED",
		)

		return
	}

	fmt.Println("mail sent")

	// ack if mail sent
	err = message.Acknowledger.Ack(message.DeliveryTag, false)

	if err != nil {
		fmt.Println("ACK failed:", err)
		return
	}

	// update status in db
	repository.DbUpdateNotificationStatus(
		context.Background(),
		(*state).OtelTracer,
		(*state).Db,
		notificationId,
		"SENT",
	)
}
