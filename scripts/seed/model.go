package main

import "database/sql"

type Seed struct {
	db *sql.DB
}

type AnimeStruct struct {
	ID        interface{} `json:"id"`
	Title     interface{} `json:"title"`
	TitleJp   interface{} `json:"title_jp"`
	Summary   interface{} `json:"summary"`
	Source    interface{} `json:"source"`
	ImageID   interface{} `json:"image_id"`
	SyoboiTid interface{} `json:"syoboi_tid"`
}

type AiringInformationStruct struct {
	AnimeID         interface{} `json:"anime_id"`
	StartDay        interface{} `json:"start_day"`
	StartMonth      interface{} `json:"start_month"`
	StartYear       interface{} `json:"start_year"`
	Year            interface{} `json:"year"`
	Season          interface{} `json:"season"`
	NumEpisodes     interface{} `json:"num_episodes"`
	EpisodeDuration interface{} `json:"episode_duration"`
	Airing          interface{} `json:"airing"`
	SyoboiTid       interface{} `json:"syoboi_tid"`
}

type StatisticStruct struct {
	AnimeID    interface{} `json:"anime_id"`
	Score      interface{} `json:"score"`
	ScoredBy   interface{} `json:"scored_by"`
	Rank       interface{} `json:"rank"`
	Popularity interface{} `json:"popularity"`
	Favorites  interface{} `json:"favorites"`
	Rating     interface{} `json:"rating"`
}

type AnimeGenreStruct struct {
	AnimeID interface{} `json:"anime_id"`
	GenreID interface{} `json:"genre_id"`
}

type AnimeStudioStruct struct {
	AnimeID  interface{} `json:"anime_id"`
	StudioID interface{} `json:"studio_id"`
}

type BroadcastTimeStruct struct {
	AnimeID   string  `json:"anime_id"`
	SyoboiTid string  `json:"syoboi_tid"`
	Time      float64 `json:"time"`
}

//
// OLD MODELS
//

// type AnimeStruct struct {
// 	ID        string  `json:"id"`
// 	Title     *string `json:"title"`
// 	TitleJp   *string `json:"title_jp"`
// 	Summary   *string `json:"summary"`
// 	Source    *string `json:"source"`
// 	ImageID   *string `json:"image_id"`
// 	SyoboiTid *int    `json:"syoboi_tid"`
// }

// type AiringInformationStruct struct {
// 	AnimeID         string  `json:"anime_id"`
// 	StartDay        *int    `json:"start_day"`
// 	StartMonth      *int    `json:"start_month"`
// 	StartYear       *int    `json:"start_year"`
// 	Year            *int    `json:"year"`
// 	Season          *string `json:"season"`
// 	NumEpisodes     *int    `json:"num_episodes"`
// 	EpisodeDuration *int    `json:"episode_duration"`
// 	Airing          *bool   `json:"airing"`
// 	SyoboiTid       *int    `json:"syoboi_tid"`
// }

// type StatisticStruct struct {
// 	AnimeID    string   `json:"anime_id"`
// 	Score      *float64 `json:"score"`
// 	ScoredBy   *int     `json:"scored_by"`
// 	Rank       *int     `json:"rank"`
// 	Popularity *int     `json:"popularity"`
// 	Favorites  *int     `json:"favorites"`
// 	Rating     *string  `json:"rating"`
// }

// type AnimeGenreStruct struct {
// 	AnimeID string `json:"anime_id"`
// 	GenreID string `json:"genre_id"`
// }

// type AnimeStudioStruct struct {
// 	AnimeID  string `json:"anime_id"`
// 	StudioID string `json:"studio_id"`
// }

// type BroadcastTimeStruct struct {
// 	SyoboiTid int     `json:"syoboi_tid"`
// 	Time      float64 `json:"time"`
// }
