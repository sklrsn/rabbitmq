package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type RConnection struct {
	conn *amqp.Connection
}

func (rc *RConnection) Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v", rabbitmqUser, rabbitmqSecret, rabbitmqHost))
	if err != nil {
		return err
	}
	rc.conn = conn

	return nil
}

func (rc *RConnection) Close() {
	if rc.conn != nil {
		rc.conn.Close()
	}
}

func (rc *RConnection) Channel() (*amqp.Channel, error) {
	return rc.conn.Channel()
}

var (
	rc = RConnection{}
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

	if len(os.Getenv("EXCHANGE_NAME")) == 0 || len(os.Getenv("EXCHANGE_NAME")) == 0 {
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(15 * time.Second)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				if rc.conn != nil {
					log.Println("Connection established ...")
					return
				}
				log.Println("Connection not established, Please check ...")
				defer cancel()
				return
			}
		}
	}()

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for rc.conn == nil {
			select {
			case <-ctx.Done():
				log.Println("context expired ....")
				return
			default:
				if err := rc.Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost); err == nil {
					log.Println("connected ...")
					return
				} else {
					log.Println(err)
				}
				time.Sleep(5 * time.Second)
			}
		}

	}(ctx, &wg)

	wg.Wait()
}

func main() {
	channel, err := rc.Channel()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := channel.ExchangeDeclare(os.Getenv("EXCHANGE_NAME"), os.Getenv("EXCHANGE_TYPE"),
		true, false, false, false, amqp.Table{}); err != nil {
		log.Fatalf("%v", err)
	}

	for i := 1; i <= 4; i++ {
		if _, err := channel.QueueDeclare(fmt.Sprintf("logs.0%v", i), true, false, false, false, amqp.Table{}); err != nil {
			log.Fatalf("%v", err)
		}
	}

	log.Println("bootstrap finished ...")
}
