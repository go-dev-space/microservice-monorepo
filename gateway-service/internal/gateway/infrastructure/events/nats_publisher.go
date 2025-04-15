package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewEventPublisher(c *nats.Conn) *NatsPublisher {
	return &NatsPublisher{
		conn: c,
	}
}

func (n *NatsPublisher) Request(topic string, data any) ([]byte, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return []byte(""), err
	}

	msg, err := n.conn.Request(topic, bs, nats.DefaultTimeout)
	if err != nil {
		return []byte(""), err
	}

	return msg.Data, nil
}
