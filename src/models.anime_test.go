package main

import "testing"

func TestGetAllAnime(t *testing.T) {
	alist := getAllAnimes()

	// Check that the length of the list of animes returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(animeList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != animeList[i].Content ||
			v.ID != animeList[i].ID ||
			v.Title != animeList[i].Title {

			t.Fail()
			break
		}
	}
}
