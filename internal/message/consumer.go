package message

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	sess      *Session
	queueName string
}

func NewConsumer(sess *Session, exchange string, routingKey string, queueName string) (*Consumer, error) {
	c, err := sess.GetChannel()
	if err != nil {
		return nil, err
	}
	q, err2 := c.QueueDeclare(queueName, true, false, false, false, nil)
	if err2 != nil {
		return nil, err2
	}
	err3 := c.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err3 != nil {
		return nil, err3
	}
	return &Consumer{
		sess:      sess,
		queueName: q.Name,
	}, nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	channel, err := c.sess.GetChannel()
	if err != nil {
		return nil, err
	}
	deliveries, err2 := channel.Consume(c.queueName, "test-tag", false, false, false, false, nil)
	if err2 != nil {
		return nil, err2
	}
	return deliveries, err
}
