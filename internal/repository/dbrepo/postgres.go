package dbrepo

import (
	"context"
	"github.com/anras5/formula-app-backend/internal/models"
	"time"
)

func (m *postgresDBRepo) SelectDrivers() ([]*models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var drivers []*models.Driver

	// --------------------------------------------------- //
	// SELECT DRIVERS
	query := `
select "ID", "NAME", "SURNAME", "COUNTRY", "START_NUMBER", created_at, updated_at, "TEAM"
from "DRIVER"
`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var driver models.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Name,
			&driver.Surname,
			&driver.Country,
			&driver.StartNumber,
			&driver.CreatedAt,
			&driver.UpdatedAt,
			&driver.Team,
		)
		if err != nil {
			return nil, err
		}

		// --------------------------------------------------- //
		// SELECT PLACES FOR DRIVER
		queryPlaces := `
select "ID", "RACE_ID", "DRIVER_ID", created_at, updated_at, "PLACE"
from "PLACE" WHERE "DRIVER_ID" = $1 
`
		rowsPlaces, errPlaces := m.DB.QueryContext(ctx, queryPlaces, driver.ID)
		if errPlaces != nil {
			return nil, errPlaces
		}

		for rowsPlaces.Next() {
			var place models.Place
			errPlaces := rowsPlaces.Scan(
				&place.ID,
				&place.RaceID,
				&place.DriverID,
				&place.CreatedAt,
				&place.UpdatedAt,
				&place.Place,
			)
			if errPlaces != nil {
				return nil, errPlaces
			}
			driver.Places = append(driver.Places, place)
		}
		if errPlaces = rowsPlaces.Err(); errPlaces != nil {
			return nil, errPlaces
		}

		rowsPlaces.Close()
		drivers = append(drivers, &driver)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return drivers, nil
}

func (m *postgresDBRepo) SelectDriver(id int) (*models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var driver models.Driver

	query := `
select "ID", "NAME", "SURNAME", "COUNTRY", "START_NUMBER", created_at, updated_at, "TEAM"
from "DRIVER"
where "ID" = $1
`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&driver.ID,
		&driver.Name,
		&driver.Surname,
		&driver.Country,
		&driver.StartNumber,
		&driver.CreatedAt,
		&driver.UpdatedAt,
		&driver.Team,
	)

	// SELECT PLACES FOR DRIVER
	queryPlaces := `
select "ID", "RACE_ID", "DRIVER_ID", created_at, updated_at, "PLACE"
from "PLACE" WHERE "DRIVER_ID" = $1 
`
	rowsPlaces, errPlaces := m.DB.QueryContext(ctx, queryPlaces, driver.ID)
	defer rowsPlaces.Close()
	if errPlaces != nil {
		return nil, errPlaces
	}

	for rowsPlaces.Next() {
		var place models.Place
		errPlaces := rowsPlaces.Scan(
			&place.ID,
			&place.RaceID,
			&place.DriverID,
			&place.CreatedAt,
			&place.UpdatedAt,
			&place.Place,
		)
		if errPlaces != nil {
			return nil, errPlaces
		}
		driver.Places = append(driver.Places, place)
	}
	if errPlaces = rowsPlaces.Err(); errPlaces != nil {
		return nil, errPlaces
	}

	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (m *postgresDBRepo) InsertDriver(driver models.Driver) (*models.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
insert into "DRIVER" ("NAME", "SURNAME", "COUNTRY", "START_NUMBER", created_at, updated_at, "TEAM")
values ($1, $2, $3, $4, $5, $6, $7) returning "ID"
`

	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	err := m.DB.QueryRowContext(ctx, stmt,
		driver.Name,
		driver.Surname,
		driver.Country,
		driver.StartNumber,
		driver.CreatedAt,
		driver.UpdatedAt,
		driver.Team,
	).Scan(&driver.ID)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (m *postgresDBRepo) UpdateDriver(driver models.Driver) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
update "DRIVER" set "NAME" = $1, "SURNAME" = $2, "COUNTRY" = $3, "START_NUMBER" = $4, "TEAM" = $5, updated_at = $6
where "ID" = $7
`
	_, err := m.DB.ExecContext(ctx, stmt,
		driver.Name,
		driver.Surname,
		driver.Country,
		driver.StartNumber,
		driver.Team,
		time.Now(),
		driver.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) DeleteDriver(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
delete from "DRIVER" where "ID" = $1
`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil

}

func (m *postgresDBRepo) SelectRaces() ([]*models.Race, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var races []*models.Race

	query := `
select "ID", "NAME", "DETAILS", "SEASON_ID", "TRACK_ID", created_at, updated_at
from "RACE"
`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var race models.Race
		err := rows.Scan(
			&race.ID,
			&race.Name,
			&race.Details,
			&race.SeasonID,
			&race.TrackID,
			&race.CreatedAt,
			&race.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		races = append(races, &race)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return races, nil
}

func (m *postgresDBRepo) InsertSeason(season models.Season) (*models.Season, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
insert into "SEASON" ("YEAR", created_at, updated_at)
values ($1, $2, $3) returning "ID"
`

	season.CreatedAt = time.Now()
	season.UpdatedAt = time.Now()
	err := m.DB.QueryRowContext(ctx, stmt,
		season.Year,
		season.CreatedAt,
		season.UpdatedAt,
	).Scan(&season.ID)
	if err != nil {
		return nil, err
	}
	return &season, nil
}
