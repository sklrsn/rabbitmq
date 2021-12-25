package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

var (
	mode                 string
	workers              int
	exchange, routingkey string
	frequency            int
	queue                string
)

var (
	rabbitmqUser, rabbitmqSecret, rabbitmqHost string
)

var (
	modeProduce = "produce"
	modeConsume = "consume"
)

var (
	rc Connection
)

func init() {
	flag.IntVar(&workers, "workers", 1, "--workers=10")
	flag.StringVar(&mode, "mode", "", "--mode=producer")
	flag.StringVar(&exchange, "exchange", "", "--exchange=logs")
	flag.StringVar(&routingkey, "routingkey", "", "--routingkey=linux")
	flag.IntVar(&frequency, "frequency", 5, "--frequency=5 (in secs)")
	flag.StringVar(&queue, "queue", "", "--queue=logs.01")

	flag.StringVar(&rabbitmqHost, "host", "", "--host=localhost")
	flag.StringVar(&rabbitmqSecret, "secret", "", "--secret=guest")
	flag.StringVar(&rabbitmqUser, "user", "", "--user=guest")

	flag.Parse()
}

func main() {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGTERM)

	if err := rc.Connect(rabbitmqUser, rabbitmqSecret, rabbitmqHost); err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("RabbitMQ connection status:%v", rc.IsConnected())

	switch mode {
	case modeProduce:
		Produce(exchange, workers)
		<-ch

	case modeConsume:
		Consume(queue, frequency, workers)
		<-ch

	default:
		os.Exit(1)
	}
}

func Produce(exchange string, workers int) {
	log.Println("starting producers ..")
	for i := 0; i < workers; i++ {
		go func(id int) {
			ch, err := rc.Channel()
			if err != nil {
				log.Fatalf("%v", err)
			}
			ticker := time.NewTicker(time.Duration(frequency) * time.Second)
			for {
				select {
				case <-ticker.C:
					message := make(map[string]interface{})
					message["timestamp"] = time.Now().String()

					bs, err := json.Marshal(message)
					if err != nil {
						log.Fatalf("%v", err)
					}
					if err := ch.Publish(exchange, routingkey, false, false,
						amqp.Publishing{
							UserId:      rabbitmqUser,
							Body:        bs,
							ContentType: "application/json"}); err != nil {
						log.Fatalf("%v", err)
					}
				}
			}
		}(i)
	}
}

func Consume(queue string, frequency, workers int) {
	for i := 1; i <= workers; i++ {
		go func(id int) {
			ch, err := rc.Channel()
			if err != nil {
				log.Fatalf("%v", err)
			}

			d, err := ch.Consume(queue, strconv.Itoa(id), true, false, false, true, amqp.Table{})
			if err != nil {
				log.Fatalf("%v", err)
			}

			for {
				select {
				case m := <-d:
					var body map[string]interface{}
					if err := json.Unmarshal(m.Body, &body); err != nil {
						log.Fatalf("%v", err)
					}
					log.Println(string(m.Body))
				}
			}
		}(i)
	}
}

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
