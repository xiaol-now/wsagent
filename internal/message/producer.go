package message

import "github.com/streadway/amqp"

type Producer struct {
	sess       *Session
	exchange   string
	routingKey string
}

func NewDirectProducer(sess *Session, exchange string, routingKey string) (*Producer, error) {
	c, err := sess.GetChannel()
	if err != nil {
		return nil, err
	}
	c.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)

	// c.Confirm(false)

	return &Producer{
		sess:       sess,
		exchange:   exchange,
		routingKey: routingKey,
	}, nil
}

func (p *Producer) Publish(body []byte) error {
	c, err := p.sess.GetChannel()
	if err != nil {
		return err
	}
	return c.Publish(p.exchange, p.routingKey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	})
}
