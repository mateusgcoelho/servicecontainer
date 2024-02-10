package services

import (
	"errors"

	"github.com/mateusgcoelho/servicecontainer/internal/models"
)

var (
	servicesInMemory []*models.ServiceModel
)

type IMemoryServiceRepository interface {
	SyncServices(services []*models.ServiceModel)
	GetServices() ([]*models.ServiceModel, error)
	GetServiceByTag(tag string) (*models.ServiceModel, error)
	validateServiceIsInitialized() error
}

type memoryServiceRepository struct { }

func NewMemoryServiceRepository() IMemoryServiceRepository {
	return memoryServiceRepository{}
}

func (r memoryServiceRepository) GetServices() ([]*models.ServiceModel, error) {
	if err := r.validateServiceIsInitialized(); err != nil {
		return nil, err
	}

	return servicesInMemory, nil
}

func (r memoryServiceRepository) GetServiceByTag(tag string) (*models.ServiceModel, error) {
	if err := r.validateServiceIsInitialized(); err != nil {
		return nil, err
	}

	for _, service := range servicesInMemory {
		if service.Tag == tag {
			return service, nil
		}
	}

	return nil, nil
}

func (r memoryServiceRepository) SyncServices(services []*models.ServiceModel) {
	servicesInMemory = services
}

func (r memoryServiceRepository) validateServiceIsInitialized() error {
	if servicesInMemory == nil {
		return errors.New("Não foi possível encontrar serviços.")
	}
	return nil
}
