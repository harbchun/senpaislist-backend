package season

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Season struct {
	DB *sql.DB
}

func (a *Season) GetSeasons() []*model.Season {
	query, err := a.DB.Prepare("SELECT * FROM seasons")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query()
	if err != nil {
		log.Fatal(err)
	}

	var seasons []*model.Season

	for res.Next() {
		var season model.Season
		err = res.Scan(&season.ID, &season.Season)

		if err != nil {
			log.Fatal(err)
		}

		seasons = append(seasons, &season)
	}

	return seasons
}
