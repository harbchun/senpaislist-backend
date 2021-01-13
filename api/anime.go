package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type CreateAnimeParams struct {
	Title           string         `json:"title" binding:"required"`
	TitleJp         string         `json:"title_jp" binding:"required"`
	ShowType        string         `json:"show_type" binding:"required"`
	Source          string         `json:"source" binding: "required"`
	BeginDate       string         `json:"begin_date" binding: "required"`
	EndDate         sql.NullString `json:"end_date"`
	Genre           []string       `json:"genre" binding: "required"`
	Season          string         `json:"season" binding: "required, oneof=Spring Summer Winter Fall"`
	Year            int64          `json:"year" binding: "required"`
	Airing          bool           `json:"airing" binding: "required"`
	CurrentStatus   string         `json:"current_status" binding: "required"`
	NumEpisodes     int64          `json:"num_episodes" binding: "required"`
	EpisodeDuration string         `json:"episode_duration" binding: "required"`
	BroadcastTime   string         `json:"broadcast_time" binding: "required"`
	NextBroadcast   sql.NullString `json:"next_broadcast"`
	Score           float64        `json:"score" binding: "required"`
	ScoredBy        int64          `json:"scored_by" binding: "required"`
	Rank            int64          `json:"rank" binding: "required"`
	Popularity      int64          `json:"popularity" binding: "required"`
	Favorites       int64          `json:"favorites" binding: "required"`
	ImageUrl        string         `json:"image_url" binding: "required"`
}

func (server *Server) createAnime(ctx *gin.Context) {

}
