package jikan

import (
	"encoding/json"
	"net/http"
)

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

// TopAnime url example: "https://api.jikan.moe/v3/top/anime/1/upcoming"
type TopAnime struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int32  `json:"request_cache_expiry"`
	Top                []struct {
		MalID     int32  `json:"mal_id"`
		Rank      int32  `json:"rank"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		ImageURL  string `json:"image_url"`
		Type      string `json:"type"`
		Episodes  int32  `json:"episodes"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		Members   int32  `json:"members"`
		Score     int32  `json:"score"`
	} `json:"top"`
}

// GetTopAnime fetches API based on page and subtype
func GetTopAnime(page int64, subtype Subtype) (anime *TopAnime, err error) {
	anime = new(TopAnime)

	// Fetch url based on page # and subtype
	url := Endpoint + "/anime/" + string(page) + "/" + string(subtype)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(anime)

	return anime, nil
}
