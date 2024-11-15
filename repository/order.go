package repository

import (
	"20241115/model"
	"database/sql"
)

type Order struct {
	Db *sql.DB
}

func InitOrderRepo(db *sql.DB) *Order {
	return &Order{db}
}

func (repo *Order) Create(order *model.Order) error {
	return nil
}
