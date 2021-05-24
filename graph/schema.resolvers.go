package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/harrisonwjs/senpaislist-backend/graph/generated"
	"github.com/harrisonwjs/senpaislist-backend/graph/model"
)

func (r *queryResolver) Anime(ctx context.Context, id string) (*model.Anime, error) {
	mockAnime := &model.Anime{
		ID:          "4",
		Title:       "Jujutsu Kaisen",
		TitleJp:     "jujutsu kaisen",
		Description: "curse bad",
		Genres:      "action",
		Year:        2020,
		ImageURL:    "google.com",
	}

	return mockAnime, nil
}

func (r *queryResolver) Animes(ctx context.Context, limit *int) ([]*model.Anime, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
