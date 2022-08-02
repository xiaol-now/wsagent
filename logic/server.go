package logic

import (
	"context"
	"net/http"

	"golang.org/x/net/websocket"
)

type WebSocketServer struct {
	ctx        context.Context
	connManage *ConnManage
}

func NewWebSocketServer(ctx context.Context) *WebSocketServer {
	// 创建一个message server

	return &WebSocketServer{ctx: ctx, connManage: NewConnManage()}
}

func (wss *WebSocketServer) HandlerConn() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		conn := NewConnection(ws)
		wss.connManage.AddConn(conn)
		// conn := NewConnection(ws)
		// 创建一个
		println("已建立一个连接", ws.LocalAddr())
		// io.Copy(conn, conn)
		conn.Forward()
	})
}

func (wss *WebSocketServer) Serve() error {
	http.Handle("/echo", wss.HandlerConn())
	return http.ListenAndServe(":12345", nil)
}
