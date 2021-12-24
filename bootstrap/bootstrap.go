package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

type RConnetion struct {
	conn *amqp.Connection
}

func (rc *RConnetion) Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v", rabbitmqUser, rabbitmqSecret, rabbitmqHost))
	if err != nil {
		return err
	}
	rc.conn = conn

	return nil
}

func (rc *RConnetion) Close() {
	if rc.conn != nil {
		rc.conn.Close()
	}
}

var (
	rc = RConnetion{}
)

func init() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqSecret := os.Getenv("RABBITMQ_PASSWORD")

	if len(rabbitmqHost) == 0 || len(rabbitmqPort) == 0 || len(rabbitmqUser) == 0 || len(rabbitmqSecret) == 0 {
		log.Println("credentials required")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for rc.conn == nil {
			if err := rc.Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost); err != nil {
				log.Println(err)
			}
		}
		log.Println("connected ...")
	}(&wg)
}

func main() {
}
