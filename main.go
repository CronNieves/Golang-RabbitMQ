package main

import (
	"Golang-RabbitMQ/consumer"
	"Golang-RabbitMQ/utils"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func main() {

	// Generate mensage
	producer()

	// Consumer
	consumer.Consumer()


}


func producer(){
	//Producer
	conn, err := amqp.Dial(utils.Config.AMQPConnectionURL)
	utils.HandleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	utils.HandleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	utils.HandleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := utils.AddTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
	body, err := json.Marshal(addTask)
	if err != nil {
		utils.HandleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)
}