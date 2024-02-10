package services

import (
	"errors"
	"fmt"
)

type ISyncServicesUsecase interface {
	Call() error
}

type syncServicesUsecase struct {
	databaseServiceRepository IDatabaseServiceRepository
	getServicesApiUsecase     IGetServicesApiUsecase
	memoryServiceRepository   IMemoryServiceRepository
}

func NewSyncServicesUsecase(
	databaseServiceRepository IDatabaseServiceRepository,
	getServicesApiUsecase IGetServicesApiUsecase,
	memoryServiceRepository IMemoryServiceRepository,
) ISyncServicesUsecase {
	return syncServicesUsecase{
		databaseServiceRepository: databaseServiceRepository,
		getServicesApiUsecase:     getServicesApiUsecase,
		memoryServiceRepository:   memoryServiceRepository,
	}
}

func (r syncServicesUsecase) Call() error {
	servicesGestor, errApi := r.getServicesApiUsecase.Call()
	if errApi != nil {
		return errApi
	}

	r.databaseServiceRepository.SyncServices(servicesGestor)

	servicesInDatabase, errDb := r.databaseServiceRepository.GetServices()
	if errDb != nil {
		return errDb
	}

	fmt.Println(fmt.Sprintf("servicesInDatabase - %d", len(servicesInDatabase)))

	if len(servicesInDatabase) == 0 && errApi != nil {
		return errors.New("Não foi possível encontrar nenhum serviço na sincronização.")
	}

	for _, i := range servicesInDatabase {
		fmt.Println(i.Id)
		fmt.Println(i.DisplayName)
		fmt.Println(i.FileName)
	}

	r.memoryServiceRepository.SyncServices(servicesInDatabase)

	return nil
}
