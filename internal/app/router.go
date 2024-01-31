package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/servicecontainer/internal/services"
)

func InitWebServer() {
	r := gin.Default()

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	services.SetupRoutes(r)

	go func() {
		if err := r.Run(serverPort); err != nil {
			panic(err)
		}
	}()
}
