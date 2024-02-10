package services

import (
	"encoding/json"
	"net/http"

	"github.com/mateusgcoelho/servicecontainer/internal/models"
)

type IApiServiceRepository interface {
	GetServicesFromGestor() ([]*models.ServiceModel, error)
}

type apiServiceRepository struct {
}

func NewApiServiceRepository() IApiServiceRepository {
	return apiServiceRepository{}
}

func (r apiServiceRepository) GetServicesFromGestor() ([]*models.ServiceModel, error) {
	response, err := http.Get("https://gestor.free.beeceptor.com")
	if err != nil {
		return nil, err
	}

	var services []*models.ServiceModel = []*models.ServiceModel{}
	err = json.NewDecoder(response.Body).Decode(&services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
