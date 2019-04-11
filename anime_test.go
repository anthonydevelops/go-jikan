package jikan_test

import (
	"testing"

	"github.com/anthonydevelops/go-jikan"
)

func TestGetTopAnime(t *testing.T) {
	_, err := jikan.GetTopAnime(1, "airing")
	if err != nil {
		t.Error(err)
	}
}

func TestGetAnimeByID(t *testing.T) {
	_, err := jikan.GetAnimeByID(1, 1)
	if err != nil {
		t.Error(err)
	}
}
