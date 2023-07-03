package models

import "time"

// swagger:model Season
type Season struct {
	// ID of the Season
	// in: int64
	ID int `json:"id"`
	// Year of the Season
	// in: int64
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Races     []Race    `json:"races"`
}

// swagger:model Driver
type Driver struct {
	// ID of the Driver
	// in: int64
	ID int `json:"id"`
	// Name of the Driver
	// in: string
	Name string `json:"name"`
	// Surname of the Driver
	// in: string
	Surname string `json:"surname"`
	// Country of the Driver
	// in: string
	Country string `json:"country"`
	// StartNumber of the Driver
	// in: int64
	StartNumber int `json:"start_number"`
	// Team of the Driver
	// in: string
	Team      string    `json:"team"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Places    []Place   `json:"places"`
}

// swagger:model Place
type Place struct {
	// ID of the Place
	// in: int64
	ID int `json:"id"`
	// Place of the Driver in a Race
	// in: int64
	Place     int       `json:"place"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	RaceID    int       `json:"race_id,omitempty"`
	DriverID  int       `json:"driver_id,omitempty"`
}

// swagger:model Race
type Race struct {
	// ID of the Race
	// in: int64
	ID int `json:"id"`
	// Name of the Race
	// in: string
	Name string `json:"name"`
	// Details of the Race
	// in: string
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	SeasonID  int       `json:"season_id,omitempty"`
	TrackID   int       `json:"track_id,omitempty"`
	Places    []Place   `json:"places,omitempty"`
}

// swagger:model Track
type Track struct {
	// ID of the Track
	// in: int64
	ID int `json:"id"`
	// Name of the Track
	// in: string
	Name string `json:"name"`
	// Details of the Track
	// in: string
	Details string `json:"details"`
	// Country of the Track
	// in: string
	Country string `json:"country"`
	// City of the Track
	// in: string
	City string `json:"city"`
	// Length of the Track
	// in: int64
	Length    int       `json:"length"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Races     []Race    `json:"races"`
}
