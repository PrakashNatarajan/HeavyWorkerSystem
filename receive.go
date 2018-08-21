package main

import (
	"fmt"
	"log"
	"time"
	"reflect"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func pageWorker(channel *amqp.Channel, queue amqp.Queue) {
	fmt.Println("Channel type", reflect.TypeOf(channel))
	fmt.Println("queue type", reflect.TypeOf(queue))

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	fmt.Println("messages type", reflect.TypeOf(msgs))

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			time.Sleep(5)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	//fmt.Println("Channel type", reflect.TypeOf(channel))
	defer channel.Close()

	listQueue, listErr := channel.QueueDeclare(
		"MMCDetailsPageQueue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(listErr, "Failed to declare a queue")

	detailQueue, detailErr := channel.QueueDeclare(
		"SMCDetailsPage", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(detailErr, "Failed to declare a queue")

	//fmt.Println("Queue type", reflect.TypeOf(detailQueue))

	pageWorker(channel, detailQueue)
	
	pageWorker(channel, listQueue)
	
}
