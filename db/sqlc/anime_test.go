package db

import (
	"context"
	"testing"

	"github.com/harrisonwjs/senpaislist-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomAnime(t *testing.T) {
	// TODO:
	// get current year and then get end dates accordingly
	// then set Airing and Current status as well
	arg := CreateAnimeParams{
		Title:           util.RandomTitle(),
		TitleJp:         util.RandomJapaneseTitle(),
		StartDay:        util.RandomDay(),
		StartMonth:      util.RandomMonth(),
		StartYear:       util.RandomYear(),
		EndDay:          0,
		EndMonth:        0,
		EndYear:         0,
		Source:          util.RandomSource(),
		Studio:          "Studio",
		Genres:          util.RandomGenres(),
		Rating:          "R",
		Description:     util.RandomDescription(),
		Season:          util.RandomSeason(),
		Year:            string(util.RandomYear()),
		NumEpisodes:     util.RandomNumEpisodes(),
		EpisodeDuration: "23 min per ep",
		Airing:          true,
		CurrentStatus:   "Currently Airing",
		NextBroadcast:   "Tue, 19 Jan 2022 05:39:18 +0900",
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

	require.Equal(t, arg.Title, anime.Title)
	require.Equal(t, arg.TitleJp, anime.TitleJp)
	require.Equal(t, arg.StartDay, anime.StartDay)
	require.Equal(t, arg.StartMonth, anime.StartMonth)
	require.Equal(t, arg.StartYear, anime.StartYear)
	require.Equal(t, arg.EndDay, anime.EndDay)
	require.Equal(t, arg.EndMonth, anime.EndMonth)
	require.Equal(t, arg.EndYear, anime.EndYear)
	require.Equal(t, arg.Source, anime.Source)
	require.Equal(t, arg.Studio, anime.Studio)
	require.Equal(t, arg.Genres, anime.Genres)
	require.Equal(t, arg.Rating, anime.Rating)
	require.Equal(t, arg.Description, anime.Description)
	require.Equal(t, arg.Season, anime.Season)
	require.Equal(t, arg.Year, anime.Year)
	require.Equal(t, arg.NumEpisodes, anime.NumEpisodes)
	require.Equal(t, arg.EpisodeDuration, anime.EpisodeDuration)
	require.Equal(t, arg.Airing, anime.Airing)
	require.Equal(t, arg.CurrentStatus, anime.CurrentStatus)
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
