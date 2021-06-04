package animesgenres

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type AnimesGenres struct {
	DB *sql.DB
}

func (a *AnimesGenres) GetAnimeGenres(id string) []*model.AnimesGenres {
	query, err := a.DB.Prepare("SELECT * FROM animes_genres WHERE anime_id=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	var animegenres []*model.AnimesGenres

	for res.Next() {
		var animegenre model.AnimesGenres
		err = res.Scan(&animegenre.AnimeID, &animegenre.Genre)

		if err != nil {
			log.Fatal(err)
		}

		animegenres = append(animegenres, &animegenre)
	}

	return animegenres
}

func (a *AnimesGenres) GetAnimesGenre(genre string) []*model.AnimesGenres {
	query, err := a.DB.Prepare("SELECT * FROM animes_genres WHERE genre=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(genre)
	if err != nil {
		log.Fatal(err)
	}

	var animegenres []*model.AnimesGenres

	for res.Next() {
		var animegenre model.AnimesGenres
		err = res.Scan(&animegenre.AnimeID, &animegenre.Genre)

		if err != nil {
			log.Fatal(err)
		}

		animegenres = append(animegenres, &animegenre)
	}

	return animegenres
}
