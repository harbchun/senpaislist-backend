package dataloader

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

const loadersKey = "dataloaders"

type Loaders struct {
	AnimesGenresLoader      AnimesGenresLoader
	StatisticLoader         StatisticLoader
	AiringInformationLoader AiringInformationLoader
	AnimesStudiosLoader     AnimesStudiosLoader
}

func DataloaderMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			AnimesGenresLoader: AnimesGenresLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([][]*model.AnimesGenres, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "$" + strconv.Itoa(i+1)
						args[i] = ids[i]
					}
					query, err := db.Prepare("SELECT * from animes_genres WHERE animes_genres.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					log.Println("SELECT * from animes_genres WHERE animes_genres.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					if err != nil {
						log.Fatal(err)
					}

					res, err := query.Query(args...)

					if err != nil {
						log.Fatal(err)
					}

					defer res.Close()

					animesGenresSlice := []*model.AnimesGenres{}
					for res.Next() {
						animesGenre := model.AnimesGenres{}
						err := res.Scan(&animesGenre.AnimeID, &animesGenre.Genre)
						if err != nil {
							panic(err)
						}
						animesGenresSlice = append(animesGenresSlice, &animesGenre)
					}

					animesGenres := make([][]*model.AnimesGenres, len(ids))

					for index, id := range ids {
						for _, genre := range animesGenresSlice {
							if genre.AnimeID == id {
								animesGenres[index] = append(animesGenres[index], genre)
							}
						}
					}

					return animesGenres, nil
				},
			},
			AnimesStudiosLoader: AnimesStudiosLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([][]*model.AnimesStudios, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "$" + strconv.Itoa(i+1)
						args[i] = ids[i]
					}
					query, err := db.Prepare("SELECT * FROM animes_studios WHERE animes_studios.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					log.Println("SELECT * FROM animes_studios WHERE animes_studios.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					if err != nil {
						log.Fatal(err)
					}

					res, err := query.Query(args...)

					if err != nil {
						log.Fatal(err)
					}

					defer res.Close()

					animesStudiosSlice := []*model.AnimesStudios{}
					for res.Next() {
						animesStudio := model.AnimesStudios{}
						err := res.Scan(&animesStudio.AnimeID, &animesStudio.Studio)
						if err != nil {
							panic(err)
						}
						animesStudiosSlice = append(animesStudiosSlice, &animesStudio)
					}

					animesStudios := make([][]*model.AnimesStudios, len(ids))

					for index, id := range ids {
						for _, studio := range animesStudiosSlice {
							if *studio.AnimeID == id {
								animesStudios[index] = append(animesStudios[index], studio)
							}
						}
					}

					return animesStudios, nil
				},
			},
			StatisticLoader: StatisticLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*model.Statistic, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "$" + strconv.Itoa(i+1)
						args[i] = ids[i]
					}
					query, err := db.Prepare("SELECT * FROM statistics WHERE statistics.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					log.Println("SELECT * FROM statistics WHERE statistics.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					if err != nil {
						log.Fatal(err)
					}

					res, err := query.Query(args...)

					if err != nil {
						log.Fatal(err)
					}

					defer res.Close()

					statisticMap := map[string]*model.Statistic{}
					for res.Next() {
						statistic := model.Statistic{}
						err := res.Scan(&statistic.AnimeID, &statistic.Score, &statistic.ScoredBy, &statistic.Rank, &statistic.Popularity, &statistic.Favorites)
						if err != nil {
							panic(err)
						}
						statisticMap[statistic.AnimeID] = &statistic
					}

					statistics := make([]*model.Statistic, len(ids))
					for i, id := range ids {
						statistics[i] = statisticMap[id]
					}

					return statistics, nil
				},
			},
			AiringInformationLoader: AiringInformationLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*model.AiringInformation, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "$" + strconv.Itoa(i+1)
						args[i] = ids[i]
					}
					query, err := db.Prepare("SELECT * FROM airing_informations WHERE airing_informations.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					log.Println("SELECT * FROM airing_informations WHERE airing_informations.anime_id IN (" + strings.Join(placeholders, ",") + ")")
					if err != nil {
						log.Fatal(err)
					}

					res, err := query.Query(args...)

					if err != nil {
						log.Fatal(err)
					}

					defer res.Close()

					AiringInformationMap := map[string]*model.AiringInformation{}
					for res.Next() {
						airingInformation := model.AiringInformation{}
						err := res.Scan(&airingInformation.AnimeID, &airingInformation.StartDay, &airingInformation.StartMonth, &airingInformation.StartYear, &airingInformation.Year, &airingInformation.Season, &airingInformation.NumEpisodes, &airingInformation.EpisodeDuration, &airingInformation.Airing, &airingInformation.SyoboiTid)
						if err != nil {
							panic(err)
						}
						AiringInformationMap[airingInformation.AnimeID] = &airingInformation
					}

					airingInformations := make([]*model.AiringInformation, len(ids))
					for i, id := range ids {
						airingInformations[i] = AiringInformationMap[id]
					}

					return airingInformations, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
