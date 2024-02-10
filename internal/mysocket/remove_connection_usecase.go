package mySocket

import (
	"github.com/gorilla/websocket"
)

type IRemoveConnectionUsecase interface {
	Call(conn *websocket.Conn)
}

type removeConnectionUsecase struct {
	connections map[*websocket.Conn]bool
}

func NewRemoveConnectionUsecase(connections map[*websocket.Conn]bool) IRemoveConnectionUsecase {
	return removeConnectionUsecase{
		connections: connections,
	}
}

func (u removeConnectionUsecase) Call(conn *websocket.Conn) {
	conn.Close()
	delete(u.connections, conn)
}
