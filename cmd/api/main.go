package main

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/mateusgcoelho/servicecontainer/internal/app"
)

var (
	db          *sql.DB
	connections map[*websocket.Conn]bool
)

func main() {
	app.InitConfigEnvs()

	db = app.InitSqliteDatabase()

	connections = make(map[*websocket.Conn]bool)
	app.InitWebServer(connections, db)

	defer db.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop
}
