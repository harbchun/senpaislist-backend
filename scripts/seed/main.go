package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "championsclub123"
	dbname   = "postgres"

	start_year = 2010
	end_year   = 2021
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var DataBucketName string
var BroadcastTimesBucketName string

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func ConnectAws() *session.Session {
	AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion = GetEnvWithKey("AWS_REGION")
	DataBucketName = GetEnvWithKey("AWS_DATA_BUCKET")
	BroadcastTimesBucketName = GetEnvWithKey("AWS_BROADCAST_INFO_BUCKET")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}

	return sess
}

func main() {
	LoadEnv()                // load environment variables
	sess := ConnectAws()     // create session & setup variable values
	s3Client := s3.New(sess) // s3 connection

	// Connect to db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	seedStruct := Seed{db}

	// seed static data (seasons)
	seedStruct.InsertSeasons()

	// broadcast info objects
	broadcastTimesObjects, err := s3Client.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: aws.String(BroadcastTimesBucketName),
		},
	)
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", BroadcastTimesBucketName, err)
	}

	// models
	animeInsertModels := []*AnimeStruct{}
	airingInformationInsertModels := []*AiringInformationStruct{}
	StatisticInsertModels := []*StatisticStruct{}
	animeGenreInsertModels := []*AnimeGenreStruct{}
	animeStudioInsertModels := []*AnimeStudioStruct{}
	broadcastTimesInsertModels := []*BroadcastTimeStruct{}
	broadcastTimesMap := make(map[string][]float64)

	// loop through all s3 objects for broadcast info
	for _, broadcastTimesObject := range broadcastTimesObjects.Contents {
		broadcastTimesObjectReqInput := &s3.GetObjectInput{
			Bucket: aws.String(BroadcastTimesBucketName),
			Key:    aws.String(*broadcastTimesObject.Key),
		}
		broadcastTimesObjectOutput, err := s3Client.GetObject(broadcastTimesObjectReqInput)
		if err != nil {
			fmt.Println(err)
		}
		defer broadcastTimesObjectOutput.Body.Close()

		// get json body from the s3 object
		broadcastTimesBody, err := ioutil.ReadAll(broadcastTimesObjectOutput.Body)
		if err != nil {
			fmt.Println(err)
		}
		broadcastTimesBodyString := fmt.Sprintf("%s", broadcastTimesBody)
		broadcastTimesMapTmp := make(map[string][]float64)
		decoder := json.NewDecoder(strings.NewReader(broadcastTimesBodyString))
		err = decoder.Decode(&broadcastTimesMap)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error while decoding \"Broadcast Times\" json string")
		}

		broadcastTimesMap = mergeFloat64SliceMap(broadcastTimesMap, broadcastTimesMapTmp)
	}

	// loop through all s3 objects for data
	seasons := []string{
		"spring",
		"summer",
		"fall",
		"winter",
	}
	for year := start_year; year < end_year+1; year++ {
		for _, season := range seasons {
			dataObjects, err := s3Client.ListObjectsV2(
				&s3.ListObjectsV2Input{
					Bucket: aws.String(DataBucketName),
					Prefix: aws.String(strconv.Itoa(year) + "/" + season),
				},
			)
			if err != nil {
				exitErrorf("Unable to list items in bucket %q, %v", DataBucketName, err)
			}

			for _, dataObject := range dataObjects.Contents {

				// get s3 object
				dataObjectReqInput := &s3.GetObjectInput{
					Bucket: aws.String(DataBucketName),
					Key:    aws.String(*dataObject.Key),
				}
				dataObjectOutput, err := s3Client.GetObject(dataObjectReqInput)
				if err != nil {
					fmt.Println(err)
				}
				defer dataObjectOutput.Body.Close()

				// get json body from the s3 object
				dataBody, err := ioutil.ReadAll(dataObjectOutput.Body)
				if err != nil {
					fmt.Println(err)
				}
				dataBodyString := fmt.Sprintf("%s", dataBody)
				var dataMap map[string]interface{}
				decoder := json.NewDecoder(strings.NewReader(dataBodyString))
				err = decoder.Decode(&dataMap)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error while decoding \"Data\" json string")
				}

				if (year == getCurrYear() && season == getCurrSeason()) ||
					isLastSeason(year, season) {
					malid := strconv.Itoa(int(dataMap["mal_id"].(float64)))
					if val, ok := broadcastTimesMap[malid]; ok {
						animeid := dataToString(dataMap["anime_id"])
						syoboitid := dataToString(dataMap["syoboi_tid"])
						broadcastTimeInsertModelSlice := buildBroadcastTimeInsertModels(animeid, syoboitid, val)
						broadcastTimesInsertModels = append(broadcastTimesInsertModels, broadcastTimeInsertModelSlice...)
					}
				}

				// ANIME
				animeInsertModel := buildAnimeInsertModel(dataMap)
				animeInsertModels = append(animeInsertModels, animeInsertModel)

				// AIRING INFORMATION
				airingInformationInsertModel := buildAiringInformationInsertModel(dataMap, dataObject.Key)
				airingInformationInsertModels = append(airingInformationInsertModels, airingInformationInsertModel)

				// STATISTICS
				StatisticInsertModel := buildStatisticInsertModel(dataMap)
				StatisticInsertModels = append(StatisticInsertModels, StatisticInsertModel)

				// GENRES
				animeGenreInsertModelsSlice := buildAnimeGenreInsertModels(dataMap)
				animeGenreInsertModels = append(animeGenreInsertModels, animeGenreInsertModelsSlice...)

				// STUDIOS
				animeStudioInsertModelsSlice := buildAnimeStudioInsertModels(dataMap)
				animeStudioInsertModels = append(animeStudioInsertModels, animeStudioInsertModelsSlice...)
			}
		}
	}

	err, errorQuery := seedStruct.InsertAnimes(animeInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}

	err, errorQuery = seedStruct.InsertAiringInformations(airingInformationInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}

	err, errorQuery = seedStruct.InsertStatistics(StatisticInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}

	err, errorQuery = seedStruct.InsertAnimeGenres(animeGenreInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}

	err, errorQuery = seedStruct.InsertAnimeStudios(animeStudioInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}

	err, errorQuery = seedStruct.InsertBroadcastTimes(broadcastTimesInsertModels)
	if err != nil {
		if errorQuery != nil {
			fmt.Println(*errorQuery)
		}
		panic(err)
	}
}
