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

// Anime url example: "https://api.jikan.moe/v3/anime/1/episodes/1"
type Anime struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int32  `json:"request_cache_expiry"`
	EpisodesLastPage   int32  `json:"episodes_last_page"`
	Episodes           []struct {
		EpisodeID     int32  `json:"episode_id"`
		Title         string `json:"title"`
		TitleJapanese string `json:"title_japanese"`
		TitleRomanji  string `json:"title_romanji"`
		Aired         struct {
			From string `json:"from"`
			To   string `json:"to"`
			Prop struct {
				From struct {
					Day   string `json:"day"`
					Month string `json:"month"`
					Year  int32  `json:"year"`
				} `json:"from"`
				To struct {
					Day   string `json:"day"`
					Month string `json:"month"`
					Year  int32  `json:"year"`
				} `json:"to"`
			} `json:"prop"`
		} `json:"aired"`
	} `json:"episodes"`
}

// GetTopAnime fetches API based on page and subtype
func GetTopAnime(page int64, subtype Subtype) (topAnime *TopAnime, err error) {
	topAnime = new(TopAnime)

	// Fetch url based on page # and subtype
	url := Endpoint + "/anime/" + string(page) + "/" + string(subtype)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(topAnime)

	return topAnime, nil
}

// GetAnime fetches API based on page and subtype
func GetAnime(id int64, page int32) (anime *Anime, err error) {
	anime = new(Anime)

	// Fetch url based on page # and subtype
	url := Endpoint + "/anime/" + string(id) + "/episodes/" + string(page)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(anime)

	return anime, nil
}
