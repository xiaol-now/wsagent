package logic

import (
	"sync"
	"wsagent/internal/message"

	guuid "github.com/google/uuid"
)

type ConnManage struct {
	conns     *sync.Map
	msgserver *message.MessageServer
}

func NewConnManage() (*ConnManage, error) {
	msgserver, err := message.NewMessageServer()
	if err != nil {
		return nil, err
	}
	return &ConnManage{conns: &sync.Map{}, msgserver: msgserver}, nil
}

func (cm *ConnManage) AddConn(conn *Connection) string {
	uuid := newUuid()
	cm.conns.Store(uuid, conn)
	return uuid
}

func (cm *ConnManage) RemoveConn(uuid string) {
	cm.conns.Delete(uuid)
}

func (cm *ConnManage) GetConn(uuid string) *Connection {
	conn, ok := cm.conns.Load(uuid)
	if !ok {
		return nil
	}
	return conn.(*Connection)
}

func newUuid() string {
	return guuid.New().String()
}
