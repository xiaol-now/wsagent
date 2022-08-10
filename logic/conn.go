package logic

import (
	"io"
	"net"
	"time"
	"wsagent/logic/codec"
)

type Connection struct {
	cc   *codec.JsonCodec
	conn net.Conn
	uuid string
}

func NewConnection(conn net.Conn) *Connection {
	cc := codec.NewJsonCodec(conn)
	c := &Connection{conn: conn, cc: cc}
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

func (c *Connection) Close() {

}
func (c *Connection) broadcastMessage() {

}

// 连接
// 发送消息
// 接收消息
// 关闭连接，要考虑客户端主动关闭连接的处理
// 心跳检测
