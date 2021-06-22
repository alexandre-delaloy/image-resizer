package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

const url = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to server. Error: %s", err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel. Error: %s", err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("resizing", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue. Error: %s", err.Error())
	}

	forever := make(chan bool)

	files, err := ioutil.ReadDir("images/uploaded")
	if err != nil {
		log.Fatalf("Failed to open dir. Error: %s", err.Error())
	}

	for _, f := range files {
		// opening the image
		file, err := os.Open("images/uploaded/" + f.Name())
		if err != nil {
			log.Fatalf("Failed to open file. Error: %s", err.Error())
		}

		img, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		time.Sleep(time.Second * 2)
		go func(img []byte, filename string) {
			// Create a message
			msg := amqp.Publishing{
				Body: img,
			}

			// Publishing
			ch.Publish("", q.Name, false, false, msg)
			fmt.Printf("%s queued...\n", filename)
		}(img, f.Name())
	}

	<-forever
}
