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

func (r *queryResolver) Anime(ctx context.Context, id string) (*model.Anime, error) {
	anime := r.AnimeController.GetAnime(id)

	return &anime, nil
}

func (r *queryResolver) Statistic(ctx context.Context, id string) (*model.Statistic, error) {
	statistic := r.StatisticController.GetStatistic(id)

	return &statistic, nil
}

func (r *queryResolver) AiringInformation(ctx context.Context, id string) (*model.AiringInformation, error) {
	airingInformation := r.AiringInformationController.GetAiringInformation(id)

	return &airingInformation, nil
}

// Anime returns generated.AnimeResolver implementation.
func (r *Resolver) Anime() generated.AnimeResolver { return &animeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type animeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
