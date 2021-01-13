package api

import (
	"github.com/gin-gonic/gin"
)

type CreateAnimeParams struct {
	Title           string   `json:"title"`
	TitleJp         string   `json:"title_jp"`
	StartDay        int64    `json:"start_day"`
	StartMonth      int64    `json:"start_month"`
	StartYear       int64    `json:"start_year"`
	EndDay          int64    `json:"end_day"`
	EndMonth        int64    `json:"end_month"`
	EndYear         int64    `json:"end_year"`
	Source          string   `json:"source"`
	Studio          string   `json:"studio"`
	Genres          []string `json:"genres"`
	Rating          string   `json:"rating"`
	Description     string   `json:"description"`
	Season          string   `json:"season"`
	Year            string   `json:"year"`
	NumEpisodes     int64    `json:"num_episodes"`
	EpisodeDuration string   `json:"episode_duration"`
	Airing          bool     `json:"airing"`
	CurrentStatus   string   `json:"current_status"`
	NextBroadcast   string   `json:"next_broadcast"`
	Score           float64  `json:"score"`
	ScoredBy        int64    `json:"scored_by"`
	Rank            int64    `json:"rank"`
	Popularity      int64    `json:"popularity"`
	Favorites       int64    `json:"favorites"`
	ImageUrl        string   `json:"image_url"`
}

func (server *Server) createAnime(ctx *gin.Context) {

}
