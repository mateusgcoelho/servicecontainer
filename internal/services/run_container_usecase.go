package services

import "github.com/mateusgcoelho/servicecontainer/internal/models"

type RunContainerUsecase interface {
	Call(service *models.ServiceModel)
}

type runContainerUsecaseImpl struct{}

func (r *runContainerUsecaseImpl) Call(service *models.ServiceModel) {

}
