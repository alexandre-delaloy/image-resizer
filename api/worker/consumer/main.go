package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"image"
	"image/jpeg"
	"strings"

	"github.com/nfnt/resize"
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		convertJpgToPng("flower.png")

		for d := range msgs {
			// create a random name for the new image
			s1 := rand.NewSource(time.Now().Unix())
			r1 := rand.New(s1)
			name := strconv.Itoa(r1.Intn(100000))
			path := "./images/resized/" + name + ".png"

			err := ioutil.WriteFile(path, d.Body, 0644)
			if err != nil {
				log.Fatal(err)
			}

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			// Deserialize file
			img, err := png.Decode(file)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			// resizing
			m := resize.Resize(512, 512, img, resize.Lanczos3)

			// create the file
			out, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}

			// write new image to file
			err = png.Encode(out, m)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s : OK\n", path)
			time.Sleep(time.Second * 5)
		}
	}()
	<-forever
}

func convertJpgToPng(file string) {
	// extension check
	filename := strings.Split(file, ".")
	switch filename[1] {
		case "jpeg":
			fmt.Println("Jpeg file")
		case "jpg":
			fmt.Println("Jpg file")
		default:
			fmt.Println("Not supported or already png")
			return
	}
	srcFileName := "./producer/test-images/" + file
	dstFileName := "./producer/test-images/" + filename[0] + ".png"

	// Decode the JPEG data. If reading from file, create a reader with
	reader, err := os.Open(srcFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//Decode from reader to image format
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Got format String", formatString)
	fmt.Println(m.Bounds())

	//Encode from image format to writer
	f, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", dstFileName, "created")
}
