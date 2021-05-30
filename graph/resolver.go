package graph

import "github.com/harrisonwjs/senpaislist-backend/graph/controller/anime"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnimeController anime.Anime
}
