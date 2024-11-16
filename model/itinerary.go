package model

type Itinerary struct {
	Id          int    `json:"-"`
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ThingsToDo  []Todo `json:"things_to_do,omitempty"`
}
