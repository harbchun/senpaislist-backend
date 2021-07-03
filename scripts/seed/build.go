package main

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func buildAnimeInsertModel(
	animeDataMap map[string]interface{},
) *AnimeStruct {
	return &AnimeStruct{
		ID:        animeDataMap["anime_id"],
		Title:     animeDataMap["title"],
		TitleJp:   animeDataMap["title_japanese"],
		Summary:   animeDataMap["synopsis"],
		Source:    animeDataMap["source"],
		ImageID:   animeDataMap["image_id"],
		SyoboiTid: animeDataMap["syoboi_tid"],
	}
}

func buildAiringInformationInsertModel(
	animeDataMap map[string]interface{},
	key *string,
) *AiringInformationStruct {
	// interfaces
	airedInterface := animeDataMap["aired"].(map[string]interface{})
	airedPropInterface := airedInterface["prop"].(map[string]interface{})
	airedPropFromInterface := airedPropInterface["from"].(map[string]interface{})

	// values
	yearStr := strings.Split(*key, "/")[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		errors.Errorf("year string to int conversion FAILED")
	}
	season := strings.Split(*key, "/")[1]

	// model creation
	return &AiringInformationStruct{
		AnimeID:         animeDataMap["anime_id"],
		StartDay:        airedPropFromInterface["day"],
		StartMonth:      airedPropFromInterface["month"],
		StartYear:       airedPropFromInterface["year"],
		Year:            &year,
		Season:          &season,
		NumEpisodes:     animeDataMap["episodes"],
		EpisodeDuration: animeDataMap["duration"],
		Airing:          animeDataMap["airing"],
		SyoboiTid:       animeDataMap["syoboi_tid"],
	}
}

func buildStatisticInsertModel(
	animeDataMap map[string]interface{},
) *StatisticStruct {
	return &StatisticStruct{
		AnimeID:    animeDataMap["anime_id"],
		Score:      animeDataMap["score"],
		ScoredBy:   animeDataMap["scored_by"],
		Rank:       animeDataMap["rank"],
		Popularity: animeDataMap["popularity"],
		Favorites:  animeDataMap["favorites"],
		Rating:     animeDataMap["rating"],
	}
}

func buildAnimeGenreInsertModels(
	animeDataMap map[string]interface{},
) []*AnimeGenreStruct {
	// array to be returned (since an anime can have multiple genres)
	thisAnimeGenreInsertModel := []*AnimeGenreStruct{}

	genresObj := reflect.ValueOf(animeDataMap["genres"])
	for i := 0; i < genresObj.Len(); i++ {
		genreObj := genresObj.Index(i).Interface().(map[string]interface{})
		animeGenreInsertModel := &AnimeGenreStruct{
			AnimeID: animeDataMap["anime_id"],
			GenreID: genreObj["name"],
		}
		thisAnimeGenreInsertModel = append(thisAnimeGenreInsertModel, animeGenreInsertModel)
	}

	return thisAnimeGenreInsertModel
}

func buildAnimeStudioInsertModels(
	animeDataMap map[string]interface{},
) []*AnimeStudioStruct {
	// array to be returned (since an anime can have multiple studios)
	thisAnimeStudioInsertModels := []*AnimeStudioStruct{}

	studiosObj := reflect.ValueOf(animeDataMap["studios"])
	for i := 0; i < studiosObj.Len(); i++ {
		studioObj := studiosObj.Index(i).Interface().(map[string]interface{})
		animeStudioInsertModel := &AnimeStudioStruct{
			AnimeID:  animeDataMap["anime_id"].(string),
			StudioID: studioObj["name"],
		}
		thisAnimeStudioInsertModels = append(thisAnimeStudioInsertModels, animeStudioInsertModel)
	}

	return thisAnimeStudioInsertModels
}

func buildBroadcastTimeInsertModels(
	id string,
	tid string,
	times []float64,
) []*BroadcastTimeStruct {
	thisBroadcastTimeInsertModels := []*BroadcastTimeStruct{}

	for _, time := range times {
		broacastTimeInsertModel := &BroadcastTimeStruct{
			AnimeID:   id,
			SyoboiTid: tid,
			Time:      time,
		}
		thisBroadcastTimeInsertModels = append(thisBroadcastTimeInsertModels, broacastTimeInsertModel)
	}

	return thisBroadcastTimeInsertModels
}
