package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {

		fmt.Printf("error connecting to rabbit %s", err.Error())
		panic(err)
	}

	defer conn.Close()

	fmt.Println("successfully connected to rabbit instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("hello world"),
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("successfully Published message to Queue")
}
