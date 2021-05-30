package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/harrisonwjs/senpaislist-backend/graph/generated"
	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

func (r *queryResolver) Anime(ctx context.Context, id string) (*model.Anime, error) {
	anime := r.AnimeController.GetAnime(id)

	return &anime, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
