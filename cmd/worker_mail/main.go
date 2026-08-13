package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/manasss0508/notifyx-go/internals/configuration"
	"github.com/manasss0508/notifyx-go/internals/queue"
	"github.com/manasss0508/notifyx-go/internals/workers"
)

func main() {
	const WORKERTYPE string = "MAIL"

	// load configuration
	state := configuration.Load()

	//
	queueName := os.Getenv("RABBITMQ_QUEUE_MAIL_NAME")

	// creating queue and consuming it
	consumer, err := queue.CreateAndConsumerQueue((*(*state).Rbmq), queueName, WORKERTYPE)
	if err != nil {
		panic(err)
	}

	fmt.Println(" email worker started and processing message")
	workers.ProcessEachMessageMail((*state), consumer)
}
