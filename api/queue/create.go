package queue

import (
	"log"
	"os"

	"github.com/blyndusk/image-resizer/api/helpers"
	"github.com/streadway/amqp"
)

func UserCreationEmitter() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.ExitOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.ExitOnError("Failed to open a channel", err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"create_user", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	helpers.ExitOnError("Failed to declare an exchange", err)

	body := helpers.BodyFrom(os.Args)
	err = ch.Publish(
		"create_user", // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helpers.ExitOnError("Failed to publish a message", err)

	log.Printf(" [x] Sent %s", body)
}

func UserCreationReceiver() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.ExitOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.ExitOnError("Failed to open a channel", err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"create_user", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	helpers.ExitOnError("Failed to declare an exchange", err)

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	helpers.ExitOnError("Failed to declare a queue", err)

	err = ch.QueueBind(
		q.Name,        // queue name
		"",            // routing key
		"create_user", // exchange
		false,
		nil,
	)
	helpers.ExitOnError("Failed to bind a queue", err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helpers.ExitOnError("Failed to register a consumer", err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*]. To exit press CTRL+C")
	<-forever
}
