package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/servicecontainer/internal/shared"
)

func handleGetServices(ctx *gin.Context) {
	var message *string = new(string)

	services, err := getServices()
	if err != nil {
		*message = fmt.Sprintf("Não foi possível buscar serviços: %s", err.Error())
		shared.SendResponse(ctx, http.StatusInternalServerError, nil, message)
		return
	}

	*message = fmt.Sprintf("Foi encontrado(s) %d registro(s)", len(services))
	shared.SendResponse(ctx, http.StatusOK, services, message)
}
