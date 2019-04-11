// Copyright (c) 2019 Anthony Campos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jikan

import (
	"encoding/json"
	"net/http"
	"strconv"
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

// AnimeNews url example: https://api.jikan.moe/v3/anime/1/news
type AnimeNews struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int32  `json:"request_cache_expiry"`
	Articles           []struct {
		URL        string `json:"url"`
		Title      string `json:"title"`
		Date       string `json:"date"`
		AuthorName string `json:"author_name"`
		AuthorURL  string `json:"author_url"`
		ForumURL   string `json:"forum_url"`
		ImageURL   string `json:"image_url"`
		Comments   int32  `json:"comments"`
		Intro      string `json:"intro"`
	}
}

// AnimePictures url example: https://api.jikan.moe/v3/anime/1/pictures
type AnimePictures struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int32  `json:"request_cache_expiry"`
	Pictures           []struct {
		Large string `json:"large"`
		Small string `json:"small"`
	}
}

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
func GetTopAnime(page int, subtype Subtype) (topAnime *TopAnime, err error) {
	topAnime = new(TopAnime)
	// s := strconv.Itoa(page)
	// Fetch url based on page # and subtype
	url := Endpoint + "/anime/" + strconv.Itoa(page) + "/" + string(subtype)
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

// GetAnimeByID fetches API based on page and subtype
func GetAnimeByID(id int, page int) (anime *Anime, err error) {
	anime = new(Anime)

	// Fetch url based on page # and subtype
	url := Endpoint + "/anime/" + strconv.Itoa(id) + "/episodes/" + strconv.Itoa(page)
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
