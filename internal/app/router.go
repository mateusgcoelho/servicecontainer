package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	mySocket "github.com/mateusgcoelho/servicecontainer/internal/mysocket"
	"github.com/mateusgcoelho/servicecontainer/internal/services"
)

func InitWebServer(connections map[*websocket.Conn]bool, db *sql.DB) {
	var (
		addConnectionUsecase    = mySocket.NewAddConnectionUsecase(connections)
		removeConnectionUsecase = mySocket.NewRemoveConnectionUsecase(connections)

		memoryServiceRepository  = services.NewMemoryServiceRepository()
		getServicesMemoryUsecase = services.NewGetServicesMemoryUsecase(memoryServiceRepository)

		apiSeviceRepository   = services.NewApiServiceRepository()
		getServicesApiUsecase = services.NewGetServicesApiUsecase(apiSeviceRepository)

		databaseServiceRepository = services.NewDatabaseServiceRepository(db)

		syncServicesUsecase = services.NewSyncServicesUsecase(
			databaseServiceRepository,
			getServicesApiUsecase,
			memoryServiceRepository,
		)
	)

	if err := syncServicesUsecase.Call(); err != nil {
		panic(err)
	}

	r := gin.Default()

	groupServices := r.Group("/services")

	groupServices.GET("/", services.HandleGetServices(getServicesMemoryUsecase))
	r.GET("/ws", handleWebSocket(addConnectionUsecase, removeConnectionUsecase))

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	go func() {
		if err := r.Run(serverPort); err != nil {
			panic(err)
		}
	}()
}
