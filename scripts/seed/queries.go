package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nullStr = "NULL"

func dataToString(data interface{}) string {

	switch d := data.(type) {
	case int:
		return strconv.Itoa(d)
	case *int:
		return strconv.Itoa(*d)
	case string:
		return "'" + strings.ReplaceAll(d, "'", "''") + "'"
	case *string:
		return "'" + strings.ReplaceAll(*d, "'", "''") + "'"
	case float64:
		return fmt.Sprintf("%.2f", d)
	case *float64:
		return fmt.Sprintf("%.2f", *d)
	case bool:
		return fmt.Sprintf("%t", d)
	case *bool:
		return fmt.Sprintf("%t", *d)
	default:
		return "NULL"

	}
}

func (s *Seed) InsertSeasons() error {
	queryString := `INSERT INTO seasons (
			season
		)
		VALUES
		(
			'winter'
		),
		(
			'spring'
		),
		(
			'summer'
		),
		(
			'fall'
		);
	`

	query, err := s.db.Prepare(queryString)
	if err != nil {
		return err
	}

	_, err = query.Exec()
	return err
}

func (s *Seed) InsertAnimes(
	animeInsertModels []*AnimeStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO animes 
			(id, title, title_jp, source, summary, image_id, syoboi_tid) 
		VALUES 
			`
	for i, animeInsertModel := range animeInsertModels {
		queryString += fmt.Sprintf(
			"(%s, %s, %s, %s, %s, %s, %s)",
			dataToString((*animeInsertModel).ID),
			dataToString((*animeInsertModel).Title),
			dataToString((*animeInsertModel).TitleJp),
			dataToString((*animeInsertModel).Source),
			dataToString((*animeInsertModel).Summary),
			dataToString((*animeInsertModel).ImageID),
			dataToString((*animeInsertModel).SyoboiTid),
		)

		if i == len(animeInsertModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}

			queryString = `INSERT INTO animes
					(id, title, title_jp, source, summary, image_id, syoboi_tid)
				VALUES
					`
		} else {
			queryString += ","
		}
	}

	return nil, nil
}

func (s *Seed) InsertAiringInformations(
	airingInformationModels []*AiringInformationStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO airing_informations 
		(anime_id, start_day, start_month, start_year, year, season, num_episodes, episode_duration, airing, syoboi_tid) 
		VALUES `
	for i, airingInformationModel := range airingInformationModels {

		queryString += fmt.Sprintf(
			"(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
			// "'"+(*airingInformationModel).AnimeID+"'",
			// *(*airingInformationModel).StartDay,
			// *(*airingInformationModel).StartMonth,
			// *(*airingInformationModel).StartYear,
			// *(*airingInformationModel).Year,
			// "'"+*(*airingInformationModel).Season+"'",
			// *(*airingInformationModel).NumEpisodes,
			// *(*airingInformationModel).EpisodeDuration,
			// *(*airingInformationModel).Airing,
			// *(*airingInformationModel).SyoboiTid,
			dataToString((*airingInformationModel).AnimeID),
			dataToString((*airingInformationModel).StartDay),
			dataToString((*airingInformationModel).StartMonth),
			dataToString((*airingInformationModel).StartYear),
			dataToString((*airingInformationModel).Year),
			dataToString((*airingInformationModel).Season),
			dataToString((*airingInformationModel).NumEpisodes),
			dataToString((*airingInformationModel).EpisodeDuration),
			dataToString((*airingInformationModel).Airing),
			dataToString((*airingInformationModel).SyoboiTid),
		)

		if i == len(airingInformationModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}
		} else {
			queryString += ","
		}
	}

	return nil, nil
}

func (s *Seed) InsertStatistics(
	statisticModels []*StatisticStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO statistics 
		(anime_id, score, scored_by, rank, popularity, favorites) 
		VALUES `
	for i, statisticModel := range statisticModels {

		queryString += fmt.Sprintf(
			"(%s, %s, %s, %s, %s, %s)",
			// "'"+(*statisticModel).AnimeID+"'",
			// *(*statisticModel).Score,
			// *(*statisticModel).ScoredBy,
			// *(*statisticModel).Rank,
			// *(*statisticModel).Popularity,
			// *(*statisticModel).Favorites,
			dataToString((*statisticModel).AnimeID),
			dataToString((*statisticModel).Score),
			dataToString((*statisticModel).ScoredBy),
			dataToString((*statisticModel).Rank),
			dataToString((*statisticModel).Popularity),
			dataToString((*statisticModel).Favorites),
		)

		if i == len(statisticModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}
		} else {
			queryString += ","
		}
	}

	return nil, nil
}

func (s *Seed) InsertAnimeGenres(
	animeGenreModels []*AnimeGenreStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO animes_genres 
		(anime_id, genre) 
		VALUES `
	for i, animeGenreModel := range animeGenreModels {

		queryString += fmt.Sprintf(
			"(%s, %s)",
			// "'"+(*animeGenreModel).AnimeID+"'",
			// "'"+(*animeGenreModel).GenreID+"'",
			dataToString((*animeGenreModel).AnimeID),
			dataToString((*animeGenreModel).GenreID),
		)

		if i == len(animeGenreModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}
		} else {
			queryString += ","
		}
	}

	return nil, nil
}

func (s *Seed) InsertAnimeStudios(
	animeStudioModels []*AnimeStudioStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO animes_studios 
		(anime_id, studio) 
		VALUES `
	for i, animeStudioModel := range animeStudioModels {

		queryString += fmt.Sprintf(
			"(%s, %s)",
			// "'"+(*animeStudioModel).AnimeID+"'",
			// "'"+(*animeStudioModel).StudioID+"'",
			dataToString((*animeStudioModel).AnimeID),
			dataToString((*animeStudioModel).StudioID),
		)

		if i == len(animeStudioModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}
		} else {
			queryString += ","
		}
	}

	return nil, nil
}

func (s *Seed) InsertBroadcastTimes(
	broadcastTimeModels []*BroadcastTimeStruct,
) (
	error,
	*string,
) {
	queryString := `INSERT INTO broadcast_times 
		(anime_id, syoboi_tid, time) 
		VALUES `
	for i, broadcastTimeModel := range broadcastTimeModels {

		queryString += fmt.Sprintf(
			"(%s, %s, %.2f)",
			(*broadcastTimeModel).AnimeID,
			(*broadcastTimeModel).SyoboiTid,
			(*broadcastTimeModel).Time,
		)

		if i == len(broadcastTimeModels)-1 {
			queryString += ";"
			query, err := s.db.Prepare(queryString)
			if err != nil {
				return err, &queryString
			}
			_, err = query.Exec()
			if err != nil {
				return err, &queryString
			}
		} else {
			queryString += ","
		}
	}

	return nil, nil
}
