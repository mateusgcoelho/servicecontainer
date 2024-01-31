package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateusgcoelho/servicecontainer/internal/app"
	"github.com/mateusgcoelho/servicecontainer/internal/shared"
)

func main() {
	app.InitConfigEnvs()
	shared.InitSqliteDatabase()
	app.SyncServicesWithGestor()
	app.InitWebServer()

	defer shared.Db.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop
}
