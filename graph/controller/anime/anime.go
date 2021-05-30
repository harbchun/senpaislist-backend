package anime

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Anime struct {
	DB *sql.DB
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

func (a *Anime) GetAnimes(id *string, title *string) []*model.Anime {
	query, err := a.DB.Prepare("SELECT * FROM animes WHERE (id=$1 OR $1 IS NULL) AND (title=$2 OR $2 IS NULL)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id, title)
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
