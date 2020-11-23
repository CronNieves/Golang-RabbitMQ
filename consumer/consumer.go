package consumer

import (
	GolangRabbitMQ "Golang-RabbitMQ"
	amqp "github.com/streadway/amqp"
)

func consumer() {
	conn, err := amqp.Dial(GolangRabbitMQ.Config.AMQPConnectionURL)
	GolangRabbitMQ.HandleError(err,"Can't connect to AMQP")
	defer conn.Close()

	amqpChanel, err := conn.Channel()
	GolangRabbitMQ.HandleError(err,"Can't created amqpChanel")

	defer amqpChanel.Close()

}

