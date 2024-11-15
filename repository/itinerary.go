package repository

import (
	"20241115/model"
	"database/sql"
)

type Itinerary struct {
	Db *sql.DB
}

func InitItineraryRepo(db *sql.DB) *Itinerary {
	return &Itinerary{Db: db}
}

func (repo *Itinerary) Get(itinerary *model.Itinerary) error {
	return nil
}
