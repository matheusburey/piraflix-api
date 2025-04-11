package model

type Movies struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Type        string    `json:"type"`
	Code        string    `json:"code"`
	About       string    `json:"about"`
	Origin      string    `json:"origin"`
	Duration    *int      `json:"duration"`
	ReleaseDate *string   `json:"release_date"`
	Trailer     *string   `json:"trailer"`
	Description *string   `json:"description"`
	Language    *string   `json:"language"`
	Views       *int      `json:"views"`
	Active      bool      `json:"active"`
	Rating      float64   `json:"rating"`
	Categories  *[]string `json:"categories"`
}
