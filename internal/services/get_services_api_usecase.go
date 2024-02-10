package services

import "github.com/mateusgcoelho/servicecontainer/internal/models"

type IGetServicesApiUsecase interface {
	Call() ([]*models.ServiceModel, error)
}

type getServicesApiUsecase struct {
	repository IApiServiceRepository
}

func NewGetServicesApiUsecase(repository IApiServiceRepository) IGetServicesApiUsecase {
	return getServicesApiUsecase{
		repository: repository,
	}
}

func (u getServicesApiUsecase) Call() ([]*models.ServiceModel, error) {
	services, err := u.repository.GetServicesFromGestor()
	if err != nil {
		return nil, err
	}
	return services, nil
}
