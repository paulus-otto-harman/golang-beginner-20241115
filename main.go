package main

import (
	"20241115/database"
	"20241115/router"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen("20241115")
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
