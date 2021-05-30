package anime

import (
	"database/sql"
	"fmt"
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

func (a *Anime) GetAnimes() {
	fmt.Println("working animes")
}
