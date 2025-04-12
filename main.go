package main

import (
	"log"
	"net/http"
	"piraflix-api/src/config"
	"piraflix-api/src/routes"
)

func main() {
	db := config.SetupDatabase()
	config.InitSchema(db)
	defer db.Close()

	// Register routes
	mux := http.NewServeMux()
	routes.RegisterMovieRoutes(mux, db)

	log.Println("Running on port 8080 🚀")
	http.ListenAndServe(":8080", mux)
}
