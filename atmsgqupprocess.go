package atmsgqupprocess

import (
	"log"
	"github.com/streadway/amqp"
)

type MsgQueue struct {
	connection *amqp.Connection
	channel *amqp.Channel
	exchange string
	queues map[string]amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectRbtMq(connName string)(*amqp.Connection) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()
	return connection
}

func createChannel(connection *amqp.Connection)(*amqp.Channel) {
	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()
	return channel
}

func createExchange(channel *amqp.Channel, exName, exType string)() {
	err = ch.ExchangeDeclare(
        exName,   // name
        exType, // type
        true,     // durable
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    failOnError(err, "Failed to declare an exchange")
}

func createQueue(channel *amqp.Channel, quName string)(amqp.Queue) {
	queue, err := channel.QueueDeclare(
		quName, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return queue
}

func (msgqueue MsgQueue)pushMsgIntoQueue(quName, messege string)(string) {
	err := msgqueue.channel.Publish(
        msgqueue.exchange, // exchange
        msgqueue.queues[quName],     // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body: []byte(messege),
        }
    )
    failOnError(err, "Failed to publish a message")
	return quName	
}

func (msgqueue MsgQueue)pullMsgFromQueue(quName string)(<-chan amqp.Delivery) {
	messages, err := msgqueue.channel.Consume(
		quName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	return messages
}

func createMessageQueue(connName, chName, exName, exType, quName string)(MsgQueue) {
	connection := connectRbtMq(connName)
	channel := createChannel(connection)
	exchange := createExchange(channel, exName, exType)
	queue := createQueue(channel, quName)
	msgqueue := MsgQueue {
		connection : connection
		channel : channel
		exchange : exchange
		queues : {quName : queue}
	}
	return msgqueue
}