package genre

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Genre struct {
	DB *sql.DB
}

func (a *Genre) GetGenres() []*model.Genre {
	query, err := a.DB.Prepare("SELECT * FROM genres")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query()
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var genres []*model.Genre

	for res.Next() {
		var genre model.Genre
		err = res.Scan(&genre.ID, &genre.Genre)

		if err != nil {
			log.Fatal(err)
		}

		genres = append(genres, &genre)
	}

	return genres
}
