package routes

import (
	"database/sql"
	"net/http"
	"piraflix-api/src/handlers"
)

func RegisterMovieRoutes(mux *http.ServeMux, db *sql.DB) {
	movieHandler := handlers.NewMovieHandler(db)

	mux.HandleFunc("/movies", movieHandler.GetMovies)
	mux.HandleFunc("/movies/{id}", movieHandler.GetMovieByID)
}
