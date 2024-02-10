package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	mySocket "github.com/mateusgcoelho/servicecontainer/internal/mysocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(
	addConnectionUsecase mySocket.IAddConnectionUsecase,
	removeConnectionUsecase mySocket.IRemoveConnectionUsecase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		addConnectionUsecase.Call(conn)
		defer removeConnectionUsecase.Call(conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
