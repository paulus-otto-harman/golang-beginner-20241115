package main

import (
	"20241115/app"
	"20241115/database"
	"20241115/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, _ := app.LoadConfiguration()

	db := database.DbOpen(config.Database)
	defer db.Close()

	r := router.NewRouter(db)

	log.Println(fmt.Sprintf("Server started on port %s", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
}
