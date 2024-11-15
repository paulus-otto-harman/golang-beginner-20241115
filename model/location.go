package model

type Location struct {
	Id          int    `json:"id,omitempty"`
	Header      string `json:"header"`
	MapUrl      string `json:"map_url"`
	Description string `json:"description"`
}
