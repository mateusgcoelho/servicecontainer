package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func InitWebServer() {
	r := gin.Default()

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	go func() {
		if err := r.Run(serverPort); err != nil {
			panic(err)
		}
	}()
}
