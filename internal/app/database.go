package app

import "github.com/mateusgcoelho/servicecontainer/internal/services"

func SyncServicesWithGestor() {
	servicesGestor, err := services.GetServicesFromGestor()
	if err != nil {
		return
	}

	services.SyncServices(servicesGestor)
}
