package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"piraflix-api/src/model"
	"piraflix-api/src/utils"
	"strconv"
	"strings"
)

type MovieHandler struct {
	db *sql.DB
}

func NewMovieHandler(db *sql.DB) *MovieHandler {
	return &MovieHandler{db: db}
}

func parseMovieRow(rows *sql.Rows) (model.MovieSummary, error) {
	var (
		rawCategories  sql.NullString
		rawReleaseDate sql.NullString
		movie          model.MovieSummary
	)

	err := rows.Scan(
		&rawCategories,
		&movie.Image,
		&movie.Type,
		&movie.Code,
		&rawReleaseDate,
		&movie.Origin,
	)
	if err != nil {
		return movie, err
	}

	if rawCategories.Valid {
		movie.Categories = strings.Split(rawCategories.String, ",")
	} else {
		movie.Categories = []string{}
	}

	movie.ReleaseDate = utils.FromNullString(rawReleaseDate)

	return movie, nil
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	offset := (page - 1) * limit

	query := `
		SELECT categories, image, type, code, release_date, origin
		FROM movies
		LIMIT $1 OFFSET $2
	`

	rows, err := h.db.Query(query, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []model.MovieSummary

	for rows.Next() {
		movie, err := parseMovieRow(rows)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":  page,
		"limit": limit,
		"data":  movies,
	})

}

func parseFullMovie(row *sql.Row) (model.Movies, error) {
	var (
		movie model.Movies
	)

	err := row.Scan(
		&movie.ID,
		&movie.Name,
		&movie.Image,
		&movie.Type,
		&movie.Code,

		&movie.About,
		&movie.Origin,
		&movie.Duration,

		&movie.ReleaseDate,
		&movie.Trailer,
		&movie.Description,
		&movie.Language,

		&movie.Views,
		&movie.Active,
	)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (h *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	idStr := parts[2] // considerando /movies/{id}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}

	row := h.db.QueryRow(`SELECT id, name, image, type, code, about, origin, duration, release_date, trailer, description, language, views, active, rating, categories FROM movies WHERE id = $1`, id)

	movie, err := parseFullMovie(row)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": movie,
	})
}
