package statistic

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Statistic struct {
	DB *sql.DB
}

func (a *Statistic) GetStatistic(id string) model.Statistic {
	query, err := a.DB.Prepare("SELECT * FROM statistics WHERE anime_id=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var statistic model.Statistic

	if res.Next() {
		err = res.Scan(&statistic.AnimeID, &statistic.Score, &statistic.ScoredBy, &statistic.Rank, &statistic.Popularity, &statistic.Favorites, &statistic.Rating)

		if err != nil {
			log.Fatal(err)
		}
	}

	return statistic
}

func (a *Statistic) GetAnimes() {
	fmt.Println("working animes")
}
