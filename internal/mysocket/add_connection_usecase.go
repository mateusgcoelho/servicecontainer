package mySocket

import (
	"github.com/gorilla/websocket"
)

type IAddConnectionUsecase interface {
	Call(conn *websocket.Conn)
}

type addConnectionUsecase struct {
	connections map[*websocket.Conn]bool
}

func NewAddConnectionUsecase(connections map[*websocket.Conn]bool) IAddConnectionUsecase {
	return addConnectionUsecase{
		connections: connections,
	}
}

func (u addConnectionUsecase) Call(conn *websocket.Conn) {
	u.connections[conn] = true
}
