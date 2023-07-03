package repository

import "github.com/anras5/formula-app-backend/internal/models"

type DatabaseRepo interface {
	SelectDrivers() ([]*models.Driver, error)
	SelectDriver(id int) (*models.Driver, error)
	InsertDriver(driver models.Driver) (*models.Driver, error)
	UpdateDriver(driver models.Driver) error
	DeleteDriver(id int) error
	SelectRaces() ([]*models.Race, error)
	InsertSeason(season models.Season) (*models.Season, error)
}
