package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/harrisonwjs/senpaislist-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomAnime(t *testing.T) {
	arg := CreateAnimeParams{
		Title:           util.RandomTitle(),
		TitleJp:         util.RandomJapaneseTitle(),
		ShowType:        "TV",
		Source:          util.RandomSource(),
		BeginDate:       "2020-10-03T00:00:00+00:00",
		EndDate:         sql.NullString{String: "", Valid: false},
		Genre:           []string{"Action", "Demons", "Supernatural", "School", "Shounen"},
		Season:          "Fall",
		Year:            2020,
		Airing:          true,
		CurrentStatus:   "Currently Airing",
		NumEpisodes:     24,
		EpisodeDuration: "23 min per ep",
		BroadcastTime:   "Saturdays at 01:25 (JST)",
		NextBroadcast:   sql.NullString{String: "Tue, 19 Jan 2021 05:39:18 +0900", Valid: true},
		Score:           8.48,
		ScoredBy:        118870,
		Rank:            106,
		Popularity:      259,
		Favorites:       8877,
		ImageUrl:        "https://cdn.myanimelist.net/images/anime/1171/109222.jpg",
	}

	// testQueries is declared in main_test.go
	anime, err := testQueries.CreateAnime(context.Background(), arg)

	require.NoError(t, err)    // check error is not nil
	require.NotEmpty(t, anime) // check the return value is not empty object

	// check inputs matches the ouput
	// columns := [21]string{
	// 	"Title",
	// 	"TitleJp",
	// 	"ShowType",
	// 	"Source",
	// 	"BeginDate",
	// 	"EndDate",
	// 	"Genre",
	// 	"Season",
	// 	"Year",
	// 	"Airing",
	// 	"CurrentStatus",
	// 	"NumEpisodes",
	// 	"EpisodeDuration",
	// 	"BroadcastTime",
	// 	"NextBroadcast",
	// 	"Score",
	// 	"ScoredBy",
	// 	"Rank",
	// 	"Popularity",
	// 	"Favorites",
	// 	"ImageUrl",
	// }
	require.Equal(t, arg.Title, anime.Title)
	require.Equal(t, arg.TitleJp, anime.TitleJp)
	require.Equal(t, arg.ShowType, anime.ShowType)
	require.Equal(t, arg.Source, anime.Source)
	require.Equal(t, arg.BeginDate, anime.BeginDate)
	require.Equal(t, arg.EndDate, anime.EndDate)
	require.Equal(t, arg.Genre, anime.Genre)
	require.Equal(t, arg.Season, anime.Season)
	require.Equal(t, arg.Year, anime.Year)
	require.Equal(t, arg.Airing, anime.Airing)
	require.Equal(t, arg.CurrentStatus, anime.CurrentStatus)
	require.Equal(t, arg.NumEpisodes, anime.NumEpisodes)
	require.Equal(t, arg.EpisodeDuration, anime.EpisodeDuration)
	require.Equal(t, arg.BroadcastTime, anime.BroadcastTime)
	require.Equal(t, arg.NextBroadcast, anime.NextBroadcast)
	require.Equal(t, arg.Score, anime.Score)
	require.Equal(t, arg.ScoredBy, anime.ScoredBy)
	require.Equal(t, arg.Rank, anime.Rank)
	require.Equal(t, arg.Popularity, anime.Popularity)
	require.Equal(t, arg.Favorites, anime.Favorites)
	require.Equal(t, arg.ImageUrl, anime.ImageUrl)

	require.NotZero(t, anime.ID)
	require.NotZero(t, anime.CreatedAt)
}

func TestCreateAnime(t *testing.T) {
	createRandomAnime(t)
}
