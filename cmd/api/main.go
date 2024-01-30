package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateusgcoelho/servicecontainer/internal/app"
)

func main() {
	app.InitConfigEnvs()
	app.InitWebServer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop
}
