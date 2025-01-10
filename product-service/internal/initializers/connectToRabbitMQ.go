package initializers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	url := os.Getenv("RABBITMQ_URL")

	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("Connected to RabbitMQ")
			return conn, nil
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, errors.New("failed to connect to RabbitMQ after 5 attempts")
}
