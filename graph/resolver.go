package graph

import (
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/airingInformation"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/anime"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/animesgenres"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/animesstudios"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/genre"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/season"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/statistic"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/year"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnimeController             anime.Anime
	StatisticController         statistic.Statistic
	AiringInformationController airingInformation.AiringInformation
	GenreController             genre.Genre
	YearController              year.Year
	SeasonController            season.Season
	AnimesGenresController      animesgenres.AnimesGenres
	AnimesStudiosController     animesstudios.AnimesStudios
}
