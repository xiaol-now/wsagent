package logic

import (
	"encoding/json"
	"io"
	"net"
	"time"
)

type Connection struct {
	conn net.Conn
	uuid string
}

func NewConnection(uuid string, conn net.Conn) *Connection {
	c := &Connection{conn: conn, uuid: uuid}
	c.connectionSuccessNotify()
	go c.heartbeat()

	return c
}

func (c *Connection) heartbeat() {
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()
	for range t.C {
		_, _ = c.conn.Write([]byte("o"))
	}
}

func (c *Connection) Forward() {
	io.Copy(c.conn, c.conn)
}

func (c *Connection) connectionSuccessNotify() {
	body, _ := json.Marshal(Message{
		Id:      c.uuid,
		Type:    CONNECTED,
		Payload: "",
	})
	_, _ = c.conn.Write(body)
}

// 连接
// 发送消息
// 接收消息
// 关闭连接，要考虑客户端主动关闭连接的处理
// 心跳检测
