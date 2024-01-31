package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/servicecontainer/internal/models"
	"github.com/mateusgcoelho/servicecontainer/internal/shared"
)

func handleGetServices(ctx *gin.Context) {
	services, err := getServices()
	if err != nil {
		shared.SendResponseWithMessage(
			ctx,
			http.StatusInternalServerError,
			nil,
			fmt.Sprintf("Não foi possível buscar serviços: %s", err.Error()),
		)
		return
	}

	shared.SendResponseWithMessage(
		ctx,
		http.StatusOK,
		services,
		fmt.Sprintf("Foi encontrado(s) %d registro(s)", len(services)),
	)
}

func handleActiveService(ctx *gin.Context) {
	service, err := getServiceByTag(ctx.Param("tag"))
	if err != nil {
		shared.SendResponseWithMessage(
			ctx,
			http.StatusInternalServerError,
			nil,
			fmt.Sprintf("Não foi possível ativar serviço: ¨%s", err.Error()),
		)
		return
	}

	if service == nil {
		shared.SendResponseWithMessage(ctx, http.StatusNotFound, nil, "Serviço não encontrado.")
		return
	}

	validationMessages := map[models.ServiceStatus]string{
		models.ServiceStarted:    "O serviço em questão já se encontra online.",
		models.ServiceInDownload: "Aguarde o download do serviço ser concluído.",
		models.ServiceStarting:   "O serviço em questão se encontrar na inicialização.",
	}

	statusToValidate := []models.ServiceStatus{
		models.ServiceStarting,
		models.ServiceInDownload,
		models.ServiceStarted,
	}

	if containsStatus(statusToValidate, service.Status) {
		shared.SendResponseWithMessage(ctx, http.StatusBadRequest, nil, validationMessages[service.Status])
		return
	}

	shared.SendResponse(ctx, http.StatusOK, service)
}

func containsStatus(slice []models.ServiceStatus, item models.ServiceStatus) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
