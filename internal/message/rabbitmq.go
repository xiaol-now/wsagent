package message

import (
	"sync/atomic"
	"time"
	"wsagent/internal/logger"

	"github.com/streadway/amqp"
)

type Session struct {
	url  string
	conn *atomic.Value
}

func NewSession(url string) *Session {
	s := &Session{
		url:  url,
		conn: &atomic.Value{},
	}
	go s.protect()
	return s
}

func (s *Session) protect() {
	for {
		conn, err := amqp.Dial(s.url)
		if err != nil {
			log.Error("dial rabbitmq error", logger.Error(err))
			time.Sleep(time.Second * 3)
			continue
		}

		s.conn.Store(conn)
		reconnect := make(chan *amqp.Error)
		conn.NotifyClose(reconnect)
		if err := <-reconnect; err != nil {
			log.Error("rabbitmq connection closed", logger.Error(err))
			s.conn = &atomic.Value{}
			continue
		}
	}
}

func (s *Session) GetConn() *amqp.Connection {
	return s.conn.Load().(*amqp.Connection)
}

func (s *Session) GetChannel() (*amqp.Channel, error) {
	return s.GetConn().Channel()
}

func (s *Session) WaitFirstConnection() {
	for {
		if s.conn.Load() != nil {
			return
		}
		time.Sleep(time.Millisecond * 3)
	}
}

func (s *Session) Close() error {
	return s.GetConn().Close()
}
