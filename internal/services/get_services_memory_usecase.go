package services

import "github.com/mateusgcoelho/servicecontainer/internal/models"

type IGetServicesMemoryUsecase interface {
	Call() ([]*models.ServiceModel, error)
}

type getServicesMemoryUsecase struct {
	repository IMemoryServiceRepository
}

func NewGetServicesMemoryUsecase(repository IMemoryServiceRepository) IGetServicesMemoryUsecase {
	return getServicesMemoryUsecase{
		repository: repository,
	}
}

func (u getServicesMemoryUsecase) Call() ([]*models.ServiceModel, error) {
	services, err := u.repository.GetServices()
	if err != nil {
		return nil, err
	}

	return services, nil
}
