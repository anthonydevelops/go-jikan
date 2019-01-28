package anime

import (
	"net/http"
)

// URL example: "https://api.jikan.moe/v3/top/anime/1/upcoming"
const (
	BASEURL = "https://api.jikan.moe/v3"
)

// TopAnime struct for top anime
type TopAnime struct {
	MalID     string `json:"mal_id"`
	Rank      int64  `json:"rank"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	ImageURL  string `json:"image_url"`
	Type      string `json:"TV"`
	Episodes  int64  `json:"episodes"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Members   int64  `json:"members"`
	Score     int64  `json:"score"`
}

// Subtype parameter definition
type Subtype string

// Subtype parameter types
const (
	Airing   Subtype = "airing"
	Upcoming Subtype = "upcoming"
	TV       Subtype = "tv"
	Movie    Subtype = "movie"
	Ova      Subtype = "ova"
	Special  Subtype = "special"
)

// GetResponse fetches API based on type, page, and subtype
func GetResponse(key string, page int64, subtype string) (*http.Response, error) {
	// This is where i'll bring the whole application together
	url := BASEURL + "/" + key + "/" + string(page) + "/" + subtype
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	return res, nil
}
