package message

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type MessageServer struct {
	sess     *Session
	producer *Producer
	consumer *Consumer
}

func NewMessageServer() (*MessageServer, error) {
	sess := NewSession(viper.GetString("rabbitmq.url"))
	exchange, routingKey, queueName := viper.GetString("rabbitmq.exchange"), viper.GetString("rabbitmq.routingKey"), viper.GetString("rabbitmq.queueName")
	producer, err := NewDirectProducer(sess, exchange, routingKey)
	if err != nil {
		return nil, err
	}
	consumer, err2 := NewConsumer(sess, exchange, routingKey, queueName)
	if err2 != nil {
		return nil, err2
	}
	return &MessageServer{
		sess:     sess,
		producer: producer,
		consumer: consumer,
	}, nil
}

func (m *MessageServer) Publish(body []byte) error {
	return m.producer.Publish(body)
}

func (m *MessageServer) Consume() (<-chan amqp.Delivery, error) {
	return m.consumer.Consume()
}
