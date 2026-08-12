package queue

import (
	"context"
	"encoding/json"
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
func (q QueueConn) ch() *amqp.Channel {
	ch, _ := (*q.conn).Channel()
	return ch
}

func (q QueueConn) Publish(ctx context.Context, otelTracer trace.Tracer, notifId uuid.UUID, channelType string) error {
	// otel
	_, span := otelTracer.Start(ctx, "pushing message to RBMQ")
	defer span.End()

	// create channel
	ch := q.ch()

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
