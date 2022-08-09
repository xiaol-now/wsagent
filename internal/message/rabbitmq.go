package message

import (
	"github.com/streadway/amqp"
	"sync/atomic"
)

type Session struct {
	url  string
	conn *atomic.Value
}

func NewSession(url string) {
}

func (s *Session) protect() {
	for {
		conn, err := amqp.Dial(s.url)
		_ = err

		s.conn.Store(conn)
		reconnect := make(chan *amqp.Error)
		conn.NotifyClose(reconnect)
		if err := <-reconnect; err != nil {
			s.conn = &atomic.Value{}

		}
	}
}
