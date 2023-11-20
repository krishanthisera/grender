package invalidate

import (
	"context"
	"time"

	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (C AMQP) Push(payload *[]string) error {
	connection, err := amqp.Dial(C.URI)
	if err != nil {
		return err
	}

	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(C.TimeOut)*time.Second)

	if cancel != nil {
		cancel()
	}

	channel, err := connection.Channel()

	if err != nil {
		return err
	}

	defer cancel()

	q, err := channel.QueueDeclare(C.QueueName, true, false, false, false, nil)

	if err != nil {
		return err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = channel.PublishWithContext(ctx, C.Exchange, q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        jsonPayload,
	})
	if err != nil {
		return err
	}
	return nil
}
