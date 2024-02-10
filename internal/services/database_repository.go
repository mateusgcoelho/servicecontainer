package services

import (
	"database/sql"
	"fmt"

	"github.com/mateusgcoelho/servicecontainer/internal/models"
)

type IDatabaseServiceRepository interface {
	GetServices() ([]*models.ServiceModel, error)
	SyncServices(services []*models.ServiceModel) error
}

type databaseServiceRepository struct {
	db *sql.DB
}

func NewDatabaseServiceRepository(db *sql.DB) IDatabaseServiceRepository {
	return databaseServiceRepository{
		db: db,
	}
}

func (r databaseServiceRepository) GetServices() ([]*models.ServiceModel, error) {
	query := `
		SELECT
			id, tag, suffixUrl,
			defaultPort, displayName,
			fileName, engineType
		FROM services
	`

	rows, err := r.db.Query(query)
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
			&service.SuffixUrl,
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

func (r databaseServiceRepository) SyncServices(services []*models.ServiceModel) error {
	for _, service := range services {
		query := `
			INSERT INTO services (
				id, tag, suffixUrl,
				defaultPort, displayName,
				fileName, engineType
			) VALUES (?, ?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING
		`

		stmt, err := r.db.Prepare(query)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println(query)

		_, err = stmt.Exec(
			service.Id,
			service.Tag,
			service.SuffixUrl,
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
