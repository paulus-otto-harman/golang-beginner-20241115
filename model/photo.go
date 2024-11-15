package model

type Photo struct {
	Id          int    `json:"-"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}
