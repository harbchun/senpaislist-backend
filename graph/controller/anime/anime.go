package anime

import (
	"database/sql"
	"log"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Anime struct {
	DB *sql.DB
}

func queryBuilder(f *model.AnimeFilterInput, j *model.AnimeSortInput) (string, string, string, string) {
	query, joinMap := filterFunc(f)

	joinOrderMap, order, group := sortFunc(j, joinMap)

	var join string

	for _, element := range joinOrderMap {
		join += element
	}

	return query, join, order, group
}

func sortFunc(f *model.AnimeSortInput, j map[string]string) (map[string]string, string, string) {
	if f != nil {
		return animeSortInput(f, j)
	}

	return j, "", ""
}

func animeSortInput(f *model.AnimeSortInput, j map[string]string) (map[string]string, string, string) {
	if len(f.Statistics) > 0 {
		for _, sort := range f.Statistics {
			if sort != nil {
				return statisticsSortInput(sort, j)
			}
		}
	}

	return j, "", ""
}

func statisticsSortInput(f *model.StatisticsSortInput, j map[string]string) (map[string]string, string, string) {
	s := ""
	g := ""
	if f.Popularity != nil {
		j["statistics"] = " INNER JOIN statistics ON animes.id = statistics.anime_id "
		if *f.Popularity == "asc" {
			s = " ORDER BY statistics.popularity ASC"
		} else {
			s = " ORDER BY statistics.popularity DESC"
		}
		g = ", statistics.popularity "
	}

	return j, s, g
}

func airingInformationsSortInput(f *model.AiringInformationsSortInput, j map[string]string) (map[string]string, string) {
	s := ""
	return j, s
}

func filterFunc(f *model.AnimeFilterInput) (string, map[string]string) {
	j := make(map[string]string)

	query, joinMap := animeFilterInput(f, j)

	return " WHERE " + query, joinMap
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
		if f.Genre.Eq != nil {
			s += " animes_genres.genre = " + "'" + *f.Genre.Eq + "'"
			j["genres"] = " INNER JOIN animes_genres ON animes.id = animes_genres.anime_id "
		}
	}
	if f.Genre != nil && f.AnimeID != nil {
		s += " AND "
	}
	if f.AnimeID != nil {
		if f.AnimeID.Eq != nil {
			s += " anime.id = " + "'" + *f.AnimeID.Eq + "'"
		}
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

func (a *Anime) GetAnimes(filter *model.AnimeFilterInput, orderBy *model.AnimeSortInput) []*model.Anime {
	s, j, b, g := queryBuilder(filter, orderBy)
	query, err := a.DB.Prepare("SELECT animes.* FROM animes " + j + s + " GROUP BY animes.id " + g + b)
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
