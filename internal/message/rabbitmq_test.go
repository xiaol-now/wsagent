package message

import (
	"testing"
)

var s = NewSession("amqp://guest:guest@192.168.3.52:5672/")

func TestNewSession(t *testing.T) {
	s.WaitFirstConnection()
	p, err := NewDirectProducer(s, "test-exchange", "test-routingKey")
	if err != nil {
		t.Error(err)
	}
	err = p.Publish([]byte("hello world"))
	if err != nil {
		t.Error(err)
	}

}

func TestConsumer(t *testing.T) {
	s.WaitFirstConnection()
	c, err := NewConsumer(s, "test-exchange", "test-routingKey", "test-queue")
	if err != nil {
		t.Error(err)
	}
	deliveries, err := c.Consume()
	if err != nil {
		t.Error(err)
	}
	for d := range deliveries {
		t.Log(string(d.Body))
		d.Ack(false)
		break
	}
}
