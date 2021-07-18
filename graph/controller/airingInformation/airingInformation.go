package airingInformation

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type AiringInformation struct {
	DB *sql.DB
}

func (a *AiringInformation) GetAiringInformation(id string) model.AiringInformation {
	query, err := a.DB.Prepare("SELECT * FROM airing_informations WHERE anime_id=$1")
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query(id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var airingInformation model.AiringInformation

	if res.Next() {
		err = res.Scan(&airingInformation.AnimeID, &airingInformation.StartDay, &airingInformation.StartMonth, &airingInformation.StartYear, &airingInformation.Year, &airingInformation.Season, &airingInformation.NumEpisodes, &airingInformation.EpisodeDuration, &airingInformation.Airing, &airingInformation.SyoboiTid)

		if err != nil {
			log.Fatal(err)
		}
	}

	return airingInformation
}

func (a *AiringInformation) GetAnimes() {
	fmt.Println("working animes")
}
