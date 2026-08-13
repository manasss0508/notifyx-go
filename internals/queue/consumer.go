package queue

import amqp "github.com/rabbitmq/amqp091-go"

func CreateAndConsumerQueue(q QueueConn, queueName string, workerType string) (<-chan amqp.Delivery, error) {
	//creating and binding queue
	queue, err := q.CreateQueueAndBind(queueName, workerType)
	if err != nil {
		return nil, err
	}

	// creating consumer on same queue
	delivery, err := q.CreateConsumer(queue.Name)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}
