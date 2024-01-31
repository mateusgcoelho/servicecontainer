package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mateusgcoelho/servicecontainer/internal/models"
	"github.com/mateusgcoelho/servicecontainer/internal/shared"
	"github.com/mateusgcoelho/servicecontainer/internal/storage"
)

func getServices() ([]*models.ServiceModel, error) {
	if storage.Services == nil {
		panic(fmt.Errorf("Não foi possível encontrar serviços."))
	}

	return storage.Services, nil
}

func GetServicesFromSqlite() ([]*models.ServiceModel, error) {
	query := `
		SELECT
			id, tag, prefixUrl,
			defaultPort, displayName,
			fileName, engineType
		FROM services
	`

	rows, err := shared.Db.Query(query)
	if err != nil {
		return nil, err
	}

	var services []*models.ServiceModel = []*models.ServiceModel{}

	for rows.Next() {
		var service *models.ServiceModel = &models.ServiceModel{
			Status: models.ServiceStoped,
		}

		if err := rows.Scan(
			&service.Id,
			&service.Tag,
			&service.PrefixUrl,
			&service.DefaultPort,
			&service.DisplayName,
			&service.FileName,
			&service.EngineType,
		); err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func getServiceByTag(tag string) (*models.ServiceModel, error) {
	if storage.Services == nil {
		return nil, fmt.Errorf("Não foi possível encontrar serviços.")
	}

	for _, service := range storage.Services {
		if service.Tag == tag {
			return service, nil
		}
	}

	return nil, nil
}

func GetServicesFromGestor() ([]*models.ServiceModel, error) {
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

func SyncServices(services []*models.ServiceModel) error {
	for _, service := range services {
		query := `
			INSERT INTO services (
				id, tag, prefixUrl,
				defaultPort, displayName,
				fileName, engineType
			) VALUES (?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING
		`

		stmt, err := shared.Db.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			service.Id,
			service.Tag,
			service.PrefixUrl,
			service.DefaultPort,
			service.DisplayName,
			service.FileName,
			service.EngineType,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
