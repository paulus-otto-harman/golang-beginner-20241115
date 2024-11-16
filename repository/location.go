package repository

import (
	"20241115/model"
	"database/sql"
)

type Location struct {
	Db *sql.DB
}

func InitLocationRepo(db *sql.DB) *Location {
	return &Location{Db: db}
}

func (repo *Location) Get(eventId int) (*model.Location, error) {
	query := `SELECT header, map_url, description
				    FROM locations
				JOIN events ON locations.destination_id = events.destination_id
				WHERE events.id = $1
			`
	var location model.Location
	repo.Db.QueryRow(query, eventId).Scan(&location.Header, &location.MapUrl, &location.Description)

	return &location, nil
}
