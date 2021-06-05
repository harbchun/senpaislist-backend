package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/harrisonwjs/senpaislist-backend/graph/generated"
	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

func (r *animeResolver) Statistic(ctx context.Context, obj *model.Anime) (*model.Statistic, error) {
	statistic := r.StatisticController.GetStatistic(obj.ID)

	return &statistic, nil
}

func (r *animeResolver) AiringInformation(ctx context.Context, obj *model.Anime) (*model.AiringInformation, error) {
	airingInformation := r.AiringInformationController.GetAiringInformation(obj.ID)

	return &airingInformation, nil
}

func (r *animeResolver) AnimeGenres(ctx context.Context, obj *model.Anime) ([]*model.AnimesGenres, error) {
	animegenres := r.AnimesGenresController.GetAnimeGenres(obj.ID)

	return animegenres, nil
}

func (r *genreResolver) AnimesGenre(ctx context.Context, obj *model.Genre) ([]*model.AnimesGenres, error) {
	animesgenre := r.AnimesGenresController.GetAnimesGenre(obj.Genre)

	return animesgenre, nil
}

func (r *queryResolver) Anime(ctx context.Context, id string) (*model.Anime, error) {
	anime := r.AnimeController.GetAnime(id)

	return &anime, nil
}

func (r *queryResolver) Animes(ctx context.Context, filter *model.AnimeFilterInput, orderBy *model.AnimeSortInput) ([]*model.Anime, error) {
	animes := r.AnimeController.GetAnimes(filter, orderBy)

	return animes, nil
}

func (r *queryResolver) Statistic(ctx context.Context, id string) (*model.Statistic, error) {
	statistic := r.StatisticController.GetStatistic(id)

	return &statistic, nil
}

func (r *queryResolver) AiringInformation(ctx context.Context, id string) (*model.AiringInformation, error) {
	airingInformation := r.AiringInformationController.GetAiringInformation(id)

	return &airingInformation, nil
}

func (r *queryResolver) Years(ctx context.Context) ([]*model.Year, error) {
	years := r.YearController.GetYears()

	return years, nil
}

func (r *queryResolver) Genres(ctx context.Context) ([]*model.Genre, error) {
	genres := r.GenreController.GetGenres()

	return genres, nil
}

func (r *queryResolver) Seasons(ctx context.Context) ([]*model.Season, error) {
	seasons := r.SeasonController.GetSeasons()

	return seasons, nil
}

func (r *queryResolver) AnimeGenres(ctx context.Context, id string) ([]*model.AnimesGenres, error) {
	animegenres := r.AnimesGenresController.GetAnimeGenres(id)

	return animegenres, nil
}

func (r *queryResolver) AnimesGenre(ctx context.Context, genre string) ([]*model.AnimesGenres, error) {
	animesgenre := r.AnimesGenresController.GetAnimesGenre(genre)

	return animesgenre, nil
}

// Anime returns generated.AnimeResolver implementation.
func (r *Resolver) Anime() generated.AnimeResolver { return &animeResolver{r} }

// Genre returns generated.GenreResolver implementation.
func (r *Resolver) Genre() generated.GenreResolver { return &genreResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type animeResolver struct{ *Resolver }
type genreResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
