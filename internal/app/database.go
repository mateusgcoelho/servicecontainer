package app

import (
	"github.com/mateusgcoelho/servicecontainer/internal/services"
	"github.com/mateusgcoelho/servicecontainer/internal/storage"
)

func SyncServicesInStorage() {
	servicesGestor, err := services.GetServicesFromGestor()
	if err != nil {
		return
	}

	services.SyncServices(servicesGestor)

	services, err := services.GetServicesFromSqlite()
	if err != nil {
		panic(err)
	}

	storage.Services = services
}
