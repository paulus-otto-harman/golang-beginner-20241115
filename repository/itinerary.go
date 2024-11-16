package repository

import (
	"20241115/model"
	"database/sql"
	"strings"
)

type Itinerary struct {
	Db *sql.DB
}

func InitItineraryRepo(db *sql.DB) *Itinerary {
	return &Itinerary{Db: db}
}

func (repo *Itinerary) Get(eventId int) ([]model.Itinerary, error) {
	const TodoListSeparator = "#@#@#"

	query := `
			WITH agenda AS (
				SELECT todos.itinerary_id, todos.name
				FROM todos
				ORDER BY todos.sequence
			)
			SELECT itineraries.id,itineraries.day, itineraries.title, itineraries.description,
			       (SELECT string_agg(name, $2) FROM agenda WHERE agenda.itinerary_id=itineraries.id)
				FROM itineraries
				JOIN events ON itineraries.destination_id = events.destination_id
				WHERE events.id = $1
				ORDER BY itineraries.day`

	rows, err := repo.Db.Query(query, eventId, TodoListSeparator)
	if err != nil {
		return nil, err
	}

	var itineraries []model.Itinerary
	for rows.Next() {
		var itinerary model.Itinerary
		var agenda sql.NullString
		if err := rows.Scan(&itinerary.Id, &itinerary.Day, &itinerary.Title, &itinerary.Description, &agenda); err != nil {
			return nil, err
		}

		if agenda.Valid {
			todos := strings.Split(agenda.String, TodoListSeparator)
			for _, todo := range todos {
				itinerary.ThingsToDo = append(itinerary.ThingsToDo, model.Todo{Name: todo})
			}
		}

		itineraries = append(itineraries, itinerary)
	}

	return itineraries, nil
}
