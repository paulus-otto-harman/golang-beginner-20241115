package database

import (
	"20241115/app"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PgConnect(dbName string, dbUser string, dbPassword string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost", dbUser, dbPassword, dbName))
}

func DbOpen(database app.DatabaseConfig) *sql.DB {
	db, err := PgConnect(database.Name, database.Username, database.Password)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
