package consumer

import (
	"Golang-RabbitMQ/utils"
	"log"
	"os"
	"encoding/json"
	amqp "github.com/streadway/amqp"
)

func Consumer() {
	conn, err := amqp.Dial(utils.Config.AMQPConnectionURL)
	utils.HandleError(err,"Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	utils.HandleError(err,"Can't created amqpChanel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add",true,false,false,false,nil)
	utils.HandleError(err,"Can't declare Queue")

	err = amqpChannel.Qos(1,0,false)
	utils.HandleError(err,"Could not configure QoS")


	messageChannel, err := amqpChannel.Consume(queue.Name,"",false,false,false,false,nil)
	utils.HandleError(err,"Can't read mensage")


	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &utils.AddTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}

