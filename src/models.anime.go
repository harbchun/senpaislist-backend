package main

import "errors"

type anime struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var animeList = []anime{
	{ID: 1, Title: "Anime 1", Content: "Anime 1 body"},
	{ID: 2, Title: "Anime 2", Content: "Anime 2 body"},
}

func getAllAnimes() []anime {
	return animeList
}

func getAnimeByID(id int) (*anime, error) {
	for _, a := range animeList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Anime not found")
}
