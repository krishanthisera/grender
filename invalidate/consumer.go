package invalidate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/krishanthisera/grender/backend"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (C *AMQP) Consumer(b backend.Backend) error {
	connection, err := amqp.Dial(C.URI)

	if err != nil {
		return err
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	q, err := channel.QueueDeclare(C.QueueName, true, false, false, false, nil)

	if err != nil {
		return err
	}

	err = channel.Qos(1, 0, false)

	if err != nil {
		return err
	}

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)

	if err != nil {
		return err
	}

	stopChan := make(chan bool)

	go func() {
		log.Printf("consumer ready, PID: %d", os.Getpid())
		for d := range msgs {
			log.Printf("invalidation request received: %s", d.Body)
			var urls []string
			err := json.Unmarshal([]byte(d.Body), &urls)
			if err != nil {
				fmt.Println(err)
			}
			for _, url := range urls {
				fmt.Println(b.Delete(string(url)))
			}

			// C.InvalidateCache(d.Body)
			if err := d.Ack(false); err != nil {
				log.Printf("error acknowledging message : %s", err)
			}
		}
	}()

	<-stopChan

	return nil

}
