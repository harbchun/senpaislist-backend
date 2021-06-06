package anime

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

type Anime struct {
	DB *sql.DB
}

func sortFunc(f *model.AnimeSortInput, j string) (string, string, string) {
	if f != nil {
		return animeSortInput(f, j)
	}

	return j, "", ""
}

func animeSortInput(f *model.AnimeSortInput, j string) (string, string, string) {
	if len(f.Statistics) > 0 {
		for _, sort := range f.Statistics {
			if sort != nil {
				return statisticsSortInput(sort, j)
			}
		}
	}

	return j, "", ""
}

func statisticsSortInput(f *model.StatisticsSortInput, j string) (string, string, string) {
	s := ""
	g := ""
	if f.Popularity != nil {
		if !strings.Contains(j, "INNER JOIN statistics ON animes.id = statistics.anime_id") {
			j += " INNER JOIN statistics ON animes.id = statistics.anime_id "
		}
		if *f.Popularity == "asc" {
			s = " ORDER BY statistics.popularity ASC"
		} else {
			s = " ORDER BY statistics.popularity DESC"
		}
		g = ", statistics.popularity "
	}

	return j, s, g
}

func airingInformationsSortInput(f *model.AiringInformationsSortInput, j string) (string, string) {
	s := ""
	return j, s
}

func filterFunc(f *model.AnimeFilterInput) (string, string) {
	j := ""

	query, join := animeFilterInput(f, j)
	if len(query) > 0 {
		query = " WHERE " + query[4:]
	}

	return query, join
}

func animeFilterInput(f *model.AnimeFilterInput, j string) (string, string) {
	s := ""
	if len(f.Or) > 0 {
		s += " AND "
		s += " ( "
		for i, filter := range f.Or {
			query, join := animeFilterInput(filter, j)
			query = query[4:]
			s += query
			j = join
			//remove first and of the sequence
			if i != len(f.Or)-1 {
				s += " OR "
			}
		}
		s += ")"
	}
	if len(f.And) > 0 {
		s += " AND "
		s += "("
		for i, filter := range f.And {
			query, join := animeFilterInput(filter, j)
			query = query[4:]
			s += query
			j = join
			if i != len(f.And)-1 {
				s += " AND "
			}
		}
		s += ")"
	}
	if len(f.AnimeGenres) > 0 {
		for _, filter := range f.AnimeGenres {
			query, join := animeGenresFilterInput(filter, j)
			s += query
			j += join
		}
	}
	if len(f.AiringInformations) > 0 {
		for _, filter := range f.AiringInformations {
			query, join := animeAiringInformationsInput(filter, j)
			s += query
			j += join
		}
	}
	if f.Title != nil {
		s += " AND animes.title = " + "'" + *f.Title.Eq + "'"
	}

	return s, j
}

func animeGenresFilterInput(f *model.AnimesGenresFilterInput, j string) (string, string) {
	s := ""
	lj := ""
	if f.Genre != nil {
		if f.Genre.Eq != nil {
			s += " AND animes_genres.genre = " + "'" + *f.Genre.Eq + "'"
			if !strings.Contains(j, "INNER JOIN animes_genres ON animes.id = animes_genres.anime_id") {
				lj = " INNER JOIN animes_genres ON animes.id = animes_genres.anime_id "
			}
		}
	}

	return s, lj
}

func animeAiringInformationsInput(f *model.AiringInformationsFilterInput, j string) (string, string) {
	s := ""
	lj := ""
	if f.Season != nil {
		s += " AND airing_informations.season = " + "'" + *f.Season.Eq + "'"
		if !strings.Contains(j, "INNER JOIN airing_informations ON animes.id = airing_informations.anime_id") {
			lj = " INNER JOIN airing_informations ON animes.id = airing_informations.anime_id "
		}
	}
	if f.Year != nil {
		s += " AND airing_informations.year = " + strconv.Itoa(*f.Year.Eq)
		if !strings.Contains(j, "INNER JOIN airing_informations ON animes.id = airing_informations.anime_id") {
			lj = " INNER JOIN airing_informations ON animes.id = airing_informations.anime_id "
		}
	}
	return s, lj
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
	where, join, group, order := "", "", "", ""
	if filter != nil {
		where, join = filterFunc(filter)
	}
	if orderBy != nil {
		join, order, group = sortFunc(orderBy, join)
	}

	log.Println("SELECT animes.* FROM animes " + join + where + " GROUP BY animes.id " + group + order)
	query, err := a.DB.Prepare("SELECT animes.* FROM animes " + join + where + " GROUP BY animes.id " + group + order)
	if err != nil {
		log.Fatal(err)
	}

	res, err := query.Query()
	defer res.Close()
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
