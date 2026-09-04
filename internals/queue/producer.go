package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/manasss0508/notifyx-go/internals/datamodels/rbmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/trace"
)

type QueueConn struct {
	conn         *amqp.Connection
	exchangeName string
	routingKeys  *map[string]string
}

// load all routing key from .env and convert it to map
func loadAllRoutingKeys() *map[string]string {
	return &map[string]string{
		"MAIL": os.Getenv("RABBITMQ_ROUTING_KEY_MAIL"),
	}
}

// NewQueueConn - create new QueueConn
func NewQueueConn() *QueueConn {
	// rabbitmq connection string
	url := os.Getenv("RABBITMQ_URL")

	// exchange name
	exchangeName := os.Getenv("RABBITMQ_EXCHANGE")

	// routing keys
	routingKeys := loadAllRoutingKeys()

	// creating connection to rabbitmq
	rbmq, err := amqp.Dial(url)
	if err != nil {
		panic("failed to connect rabbitmq")
	}

	return &QueueConn{
		rbmq,
		exchangeName,
		routingKeys,
	}
}

// ch - create channel on rbmq connection
func (q QueueConn) ch() (*amqp.Channel, error) {
	ch, err := (*q.conn).Channel()
	return ch, err
}

// publish message to "notifcation" exchange
func (q QueueConn) Publish(ctx context.Context, otelTracer trace.Tracer, notifId uuid.UUID, channelType string) error {
	// otel
	_, span := otelTracer.Start(ctx, "pushing message to RBMQ")
	defer span.End()

	// create channel
	ch, err := q.ch()
	if err != nil {
		return err
	}

	// create message
	msg := rbmq.NotifMsg{
		NotficationId: notifId,
	}

	// serializing message struct -> json bytes
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// publish message to exchange
	routingKey, _ := (*q.routingKeys)[channelType] // getting routing key

	return ch.PublishWithContext( // publishing message to exchange
		context.Background(),
		q.exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body: msgBytes,
		},
	)
}

// create queue and bind it "notification" exchange
// queue will be durable, non-exclusive, no auto-delete
func (q QueueConn) CreateQueueAndBind(
	queueName string,
	workerType string,
) (amqp.Queue, error) {
	// create channel of checking queue exits
	channel1, err := q.ch()
	if err != nil {
		return amqp.Queue{}, err
	}

	// check if queue exits
	queue_exists, err := channel1.QueueDeclarePassive(
		queueName,
		true,  // durable queue
		false, // no auto delete, if last connection ends queue will be not deleted
		false, // non exclusive, so other connection can access it
		false, // no fire and forget
		nil,
	)

	// if not exits create and bind
	if err != nil {
		// creating channel to create queue
		channel2, err := q.ch()
		if err != nil {
			return amqp.Queue{}, err
		}

		// creating queue
		queue, err := channel2.QueueDeclare(
			queueName,
			true,  // durable queue
			false, // no auto delete, if last connection ends queue will be not deleted
			false, // non exclusive, so other connection can access it
			false, // no fire and forget
			nil,
		)
		// routing key
		val, ok := (*q.routingKeys)[workerType]
		if !ok {
			return amqp.Queue{}, fmt.Errorf("routing key not exits for workers : %s", workerType)
		}

		// binding queue to "notification" exchange
		err = channel2.QueueBind(
			queue.Name,
			val, // routing key
			q.exchangeName,
			false,
			nil,
		)

		// binding fails then delete queue
		if err != nil {
			//deleting queue
			_, err1 := channel2.QueueDelete(queue.Name, false, false, false)
			if err1 != nil {
				// queue created, binding failed, and deleting queue also failed
				return amqp.Queue{}, err1
			}

			// queue created, binding failed, queue deleted
			return amqp.Queue{}, err
		}

		// queue created
		return queue, nil

	}

	// queue already exits
	return queue_exists, nil
}

// it will create a consumer on queue and return channel of delivery
func (q QueueConn) CreateConsumer(queueName string) (<-chan amqp.Delivery, error) {
	// creating channel
	channel, err := q.ch()
	if err != nil {
		return nil, err
	}

	notifyClose := channel.NotifyClose(make(chan *amqp.Error))

	go func() {
		err := <-notifyClose

		fmt.Println("========== CHANNEL CLOSED ==========")

		if err != nil {
			fmt.Printf("RabbitMQ error: %+v\n", err)
			fmt.Println("Code:", err.Code)
			fmt.Println("Reason:", err.Reason)
		} else {
			fmt.Println("RabbitMQ closed channel without error")
		}
	}()

	// creating consumer
	consumer, err := channel.Consume(
		queueName,
		"",
		false, // no auto ack
		true,  // exclusive, until this consumer is alive no other can consume queue
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
