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

type Connection struct {
	conn *amqp.Connection
}

func (rc *Connection) Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v", rabbitmqUser, rabbitmqSecret, rabbitmqHost))
	if err != nil {
		return err
	}
	rc.conn = conn

	return nil
}

func (rc *Connection) IsConnected() bool {
	return rc.conn != nil
}

func (rc *Connection) Close() {
	if rc.conn != nil {
		rc.conn.Close()
	}
}

func (rc *Connection) Channel() (*amqp.Channel, error) {
	return rc.conn.Channel()
}

var (
	rc = Connection{}
)

func init() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqSecret := os.Getenv("RABBITMQ_PASSWORD")

	if len(rabbitmqHost) == 0 || len(rabbitmqUser) == 0 || len(rabbitmqSecret) == 0 {
		log.Println("credentials required")
		os.Exit(1)
	}

	if len(os.Getenv("EXCHANGE_NAME")) == 0 || len(os.Getenv("EXCHANGE_TYPE")) == 0 {
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
				if rc.IsConnected() {
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
		for !rc.IsConnected() {
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
	defer rc.Close()

	channel, err := rc.Channel()
	if err != nil {
		log.Fatalf("%v", err)
	}

	//unrouted
	if err := channel.ExchangeDeclare(fmt.Sprintf("%v.unrouted", os.Getenv("EXCHANGE_NAME")),
		"topic", true, false, false, false, amqp.Table{}); err != nil {
		log.Fatalf("%v", err)
	}
	if queue, err := channel.QueueDeclare(fmt.Sprintf("%v.unrouted", os.Getenv("EXCHANGE_NAME")), true,
		false, false, false, amqp.Table{"x-queue-type": os.Getenv("QUEUE_TYPE")}); err == nil {
		err = channel.QueueBind(queue.Name, "#",
			fmt.Sprintf("%v.unrouted", os.Getenv("EXCHANGE_NAME")), false, amqp.Table{})
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	//deadletter
	if err := channel.ExchangeDeclare(fmt.Sprintf("%v.deadletter", os.Getenv("EXCHANGE_NAME")),
		"topic", true, false, false, false, amqp.Table{}); err != nil {
		log.Fatalf("%v", err)
	}
	if queue, err := channel.QueueDeclare(fmt.Sprintf("%v.deadletter", os.Getenv("EXCHANGE_NAME")), true,
		false, false, false, amqp.Table{"x-queue-type": os.Getenv("QUEUE_TYPE")}); err == nil {
		err = channel.QueueBind(queue.Name, "#",
			fmt.Sprintf("%v.deadletter", os.Getenv("EXCHANGE_NAME")), false, amqp.Table{})
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	//logs
	if err := channel.ExchangeDeclare(os.Getenv("EXCHANGE_NAME"), os.Getenv("EXCHANGE_TYPE"),
		true, false, false, false, amqp.Table{
			"alternate-exchange": fmt.Sprintf("%v.unrouted", os.Getenv("EXCHANGE_NAME"))}); err != nil {
		log.Fatalf("%v", err)
	}
	switch os.Getenv("EXCHANGE_TYPE") {
	case "topic":
		for idx := 1; idx <= 4; idx++ {
			switch idx {
			case 1:
				logQueues(channel, idx, fmt.Sprintf("%v.linux.*", os.Getenv("BINDING_KEY")))
			case 2:
				logQueues(channel, idx, fmt.Sprintf("%v.windows.*", os.Getenv("BINDING_KEY")))
			case 3:
				logQueues(channel, idx, fmt.Sprintf("%v.mac.*", os.Getenv("BINDING_KEY")))
			case 4:
				logQueues(channel, idx, fmt.Sprintf("%v.others.*", os.Getenv("BINDING_KEY")))
			}
		}
	case "direct", "fanout", "headers", "x-consistent-hash":
		for idx := 1; idx <= 4; idx++ {
			switch idx {
			case 1:
				logQueues(channel, idx, fmt.Sprintf("%v", os.Getenv("BINDING_KEY")))
			case 2:
				logQueues(channel, idx, fmt.Sprintf("%v", os.Getenv("BINDING_KEY")))
			case 3:
				logQueues(channel, idx, fmt.Sprintf("%v", os.Getenv("BINDING_KEY")))
			case 4:
				logQueues(channel, idx, fmt.Sprintf("%v", os.Getenv("BINDING_KEY")))
			}
		}
	}
	log.Println("bootstrap finished ...")
}

func logQueues(channel *amqp.Channel, idx int, bKey string) {
	if queue, err := channel.QueueDeclare(fmt.Sprintf("logs.0%v", idx), true, false, false,
		false, amqp.Table{
			"x-dead-letter-exchange": fmt.Sprintf("%v.deadletter", os.Getenv("EXCHANGE_NAME")),
			"x-queue-type":           os.Getenv("QUEUE_TYPE")}); err == nil {
		if err := channel.QueueBind(queue.Name, bKey,
			os.Getenv("EXCHANGE_NAME"), false, amqp.Table{}); err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		log.Fatalf("%v", err)
	}
}
