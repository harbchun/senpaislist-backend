// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Anime struct {
	Title           string         `json:"title"`
	TitleJp         string         `json:"title_jp"`
	ShowType        string         `json:"show_type"`
	Source          string         `json:"source"`
	BeginDate       string         `json:"begin_date"`
	EndDate         sql.NullString `json:"end_date"`
	Genre           []string       `json:"genre"`
	Season          string         `json:"season"`
	Year            int64          `json:"year"`
	Airing          bool           `json:"airing"`
	CurrentStatus   string         `json:"current_status"`
	NumEpisodes     int64          `json:"num_episodes"`
	EpisodeDuration string         `json:"episode_duration"`
	BroadcastTime   string         `json:"broadcast_time"`
	NextBroadcast   sql.NullString `json:"next_broadcast"`
	Score           float64        `json:"score"`
	ScoredBy        int64          `json:"scored_by"`
	Rank            int64          `json:"rank"`
	Popularity      int64          `json:"popularity"`
	Favorites       int64          `json:"favorites"`
	ImageUrl        string         `json:"image_url"`
	ID              int64          `json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
}
