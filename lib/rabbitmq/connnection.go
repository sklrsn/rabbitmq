package rabbitmq

import (
	"fmt"

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
