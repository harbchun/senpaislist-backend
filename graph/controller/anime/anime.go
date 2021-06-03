package anime

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Anime struct {
	DB *sql.DB
}

func filterFunc(f *model.AnimeFilterInput) (string, string) {
	j := make(map[string]string)

	query, joinMap := animeFilterInput(f, j)

	var join string

	for _, element := range joinMap {
		join += element
	}

	return " WHERE " + query, join
}

func animeFilterInput(f *model.AnimeFilterInput, j map[string]string) (string, map[string]string) {
	s := ""

	if len(f.Or) > 0 {
		s += "("
		for i, filter := range f.Or {
			if i != 0 {
				s += " OR "
			}
			query, join := animeFilterInput(filter, j)
			s += query
			j = join
		}
		s += ")"
	}
	if len(f.And) > 0 {
		s += "("
		for i, filter := range f.And {
			if i != 0 {
				s += " AND "
			}
			query, join := animeFilterInput(filter, j)
			s += query
			j = join
		}
		s += ")"
	}
	if len(f.AnimeGenres) > 0 {
		for _, filter := range f.AnimeGenres {
			query, join := animeGenresFilterInput(filter, j)
			s += query
			j = join
		}
	}

	return s, j
}

func animeGenresFilterInput(f *model.AnimesGenresFilterInput, j map[string]string) (string, map[string]string) {
	s := ""
	if f.Genre != nil {
		s += " animes_genres.genre = " + "'" + *f.Genre + "'"
		j["genres"] = " INNER JOIN animes_genres ON animes.id = animes_genres.anime_id "
	}
	if f.Genre != nil && f.AnimeID != nil {
		s += " AND "
	}
	if f.AnimeID != nil {
		s += " anime.id = " + "'" + *f.AnimeID + "'"
	}

	return s, j
}

func (a *Anime) GetAnime(id string) model.Anime {
	query, err := a.DB.Prepare("SELECT * FROM animes WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	var anime model.Anime

	if res.Next() {
		err = res.Scan(&anime.ID, &anime.Title, &anime.TitleJp, &anime.Tid, &anime.Source, &anime.Studio, &anime.Summary, &anime.ImageURL)

		if err != nil {
			log.Fatal(err)
		}
	}

	return anime
}

func (a *Anime) GetAnimes(filter *model.AnimeFilterInput) []*model.Anime {
	s, j := filterFunc(filter)

	query, err := a.DB.Prepare("SELECT animes.* FROM animes " + j + s + " GROUP BY animes.id")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query()
	if err != nil {
		log.Fatal(err)
	}

	var animes []*model.Anime

	for res.Next() {
		var anime model.Anime
		err = res.Scan(&anime.ID, &anime.Title, &anime.TitleJp, &anime.Tid, &anime.Source, &anime.Studio, &anime.Summary, &anime.ImageURL)

		if err != nil {
			log.Fatal(err)
		}

		animes = append(animes, &anime)
	}

	return animes
}
