package services

import "github.com/mateusgcoelho/servicecontainer/internal/models"

type IGetServicesDatabaseUsecase interface {
	Call() ([]*models.ServiceModel, error)
}

type getServicesDatabaseUsecase struct {
	repository IDatabaseServiceRepository
}

func NewGetServicesDatabaseUsecase(repository IDatabaseServiceRepository) IGetServicesDatabaseUsecase {
	return getServicesDatabaseUsecase{
		repository: repository,
	}
}

func (u getServicesDatabaseUsecase) Call() ([]*models.ServiceModel, error) {
	services, err := u.repository.GetServices()
	if err != nil {
		return nil, err
	}

	return services, nil
}
