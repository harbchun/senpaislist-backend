package animesstudios

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type AnimesStudios struct {
	DB *sql.DB
}

func (a *AnimesStudios) GetAnimeStudios(id string) []*model.AnimesStudios {
	query, err := a.DB.Prepare("SELECT * FROM animes_studios WHERE anime_id=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var animestudios []*model.AnimesStudios

	for res.Next() {
		var animestudio model.AnimesStudios
		err = res.Scan(&animestudio.AnimeID, &animestudio.Studio)

		if err != nil {
			log.Fatal(err)
		}

		animestudios = append(animestudios, &animestudio)
	}

	return animestudios
}
