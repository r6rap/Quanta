package main

import (
	"log"
	"net/http"

	"github.com/r6rap/Quanta/cmd/routes"
	"github.com/r6rap/Quanta/pkg/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect db")
	}
	defer db.Close()

	routes := routes.NewRoutes(db)

	r := routes.InitRoutes()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
