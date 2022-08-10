package logic

import (
	"context"
	"golang.org/x/net/websocket"
	"net/http"
)

type WebSocketServer struct {
	ctx        context.Context
	connManage *ConnManage
}

func NewWebSocketServer(ctx context.Context) (*WebSocketServer, error) {
	// 创建一个message server
	connManage, err := NewConnManage()
	if err != nil {
		return nil, err
	}
	return &WebSocketServer{ctx: ctx, connManage: connManage}, nil
}

func (ws *WebSocketServer) HandlerConn() websocket.Handler {
	return func(wsconn *websocket.Conn) {
		conn := NewConnection(wsconn)
		defer conn.Close()
		ws.connManage.AddConn(conn)

		//defer func() {
		//	println("连接已断开")
		//}()
		//
		//conn := NewConnection("todo123", ws)
		//ws.connManage.AddConn(conn)
		//// conn := NewConnection(ws)
		//// 创建一个
		//println("已建立一个连接", ws.LocalAddr())
		//// io.Copy(conn, conn)
		//conn.Forward()
	}
}

func (ws *WebSocketServer) Serve() error {
	http.Handle("/echo", ws.HandlerConn())
	return http.ListenAndServe(":12345", nil)
}
