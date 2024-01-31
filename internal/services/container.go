package services

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/mateusgcoelho/servicecontainer/internal/models"
)

var (
	mapEnginesPaths map[models.EngineType]string = map[models.EngineType]string{
		models.EngineJava8:  "/engines/java8/bin",
		models.EngineNodeJs: "/engines/nodejs/bin",
	}
	mapEnginesArgs map[models.EngineType][]string = map[models.EngineType][]string{
		models.EngineJava8:  {"-jar"},
		models.EngineNodeJs: {""},
	}
	mapEnginesExtensions map[models.EngineType]string = map[models.EngineType]string{
		models.EngineJava8:  ".jar",
		models.EngineNodeJs: ".js",
	}
	mutexService sync.Mutex
)

func RunContainerToService(service *models.ServiceModel) (*models.ServiceModel, error) {
	if service.FileName == "" {
		return nil, fmt.Errorf("Não foi possível encontrar nome de arquivo.")
	}

	go upProcessService(service)

	return nil, nil
}

func upProcessService(service *models.ServiceModel) {
	mutexService.Lock()
	defer mutexService.Unlock()

	cmd := mountCommandEngine(service)

	if err := cmd.Run(); err != nil {
		service.Status = models.ServiceStoped
		// notificar no websocket o novo status
		return
	}
}

func mountCommandEngine(service *models.ServiceModel) *exec.Cmd {
	enginePath := mapEnginesPaths[service.EngineType]
	engineArgs := strings.Join(mapEnginesArgs[service.EngineType], "")
	fileWithExtension := fmt.Sprintf(
		"%s%s",
		service.FileName,
		mapEnginesExtensions[service.EngineType],
	)

	return exec.Command(enginePath, engineArgs, fileWithExtension)
}
