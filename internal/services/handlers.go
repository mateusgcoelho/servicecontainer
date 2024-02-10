package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/servicecontainer/internal/models"
	"github.com/mateusgcoelho/servicecontainer/internal/shared"
)

var (
	validationMessages = map[models.ServiceStatus]string{
		models.ServiceStarted:    "O serviço em questão já se encontra online.",
		models.ServiceInDownload: "Aguarde o download do serviço ser concluído.",
		models.ServiceStarting:   "O serviço em questão se encontrar na inicialização.",
	}

	statusToValidate = []models.ServiceStatus{
		models.ServiceStarting,
		models.ServiceInDownload,
		models.ServiceStarted,
	}
)

func HandleGetServices(getServicesMemoryUsecase IGetServicesMemoryUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		services, err := getServicesMemoryUsecase.Call()
		if err != nil {
			shared.SendResponseWithMessage(ctx, http.StatusInternalServerError, nil, err.Error())
			return
		}

		shared.SendResponseWithMessage(ctx, http.StatusOK, services, fmt.Sprintf("Serviços encontrados."))
	}
}
