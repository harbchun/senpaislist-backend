import datetime
import json
import glob
import urllib3
import xmltodict

# next air date xml file
xml_url = "http://cal.syoboi.jp/proginfo.xml"
http = urllib3.PoolManager()
response = http.request('GET', xml_url)
data = xmltodict.parse(response.data)
nextAirDateDict = {}
for item in data['rss']['channel']['item']:
    nextAirDateDict[item['title']] = item['pubDate']

seasons = ['spring', 'summer', 'fall', 'winter']

# allAnime = []
queryString = """CREATE TABLE "anime" (
    "title" varchar NOT NULL,
    "title_jp" varchar NOT NULL,
    "start_day" bigint NOT NULL,
    "start_month" bigint NOT NULL,
    "start_year" bigint NOT NULL,
    "end_day" bigint NOT NULL,
    "end_month" bigint NOT NULL,
    "end_year" bigint NOT NULL,
    "source" varchar NOT NULL,
    "studio" varchar NOT NULL,
    "genres" varchar[] NOT NULL,
    "rating" varchar NOT NULL,
    "description" varchar NOT NULL,
    "season" varchar NOT NULL,
    "year" bigint NOT NULL,
    "num_episodes" bigint NOT NULL,
    "episode_duration" varchar NOT NULL,
    "airing" boolean NOT NULL,
    "current_status" varchar NOT NULL,
	"broadcast_time" varchar NOT NULL,
    "next_broadcast" varchar NOT NULL,
    "score" float NOT NULL,
    "scored_by" bigint NOT NULL,
    "rank" bigint NOT NULL,
    "popularity" bigint NOT NULL,
    "favorites" bigint NOT NULL,
    "image_url" varchar NOT NULL,
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
); \n
INSERT INTO anime VALUES \n\t"""
for currYear in range(2010, 2022):
    for season in seasons:
        animeFiles = glob.glob('../db/data/'+str(currYear)+str(season)+'/*.json')
        for animeFile in animeFiles:

            with open(animeFile, 'r') as jsonfile:
                animeData = json.load(jsonfile)

                title = animeData.get("title", 'N/A')
                title = 'N/A' if not title else title
                if title and title != 'N/A':
                #     title = title.replace('\"', '')
                    title = title.replace("\'", "\'\'")
                title_jp = animeData.get("title_japanese", 'N/A')
                title_jp = 'N/A' if not title_jp else title_jp
                # if title_jp and title_jp != 'N/A':
                    

                start_day = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("day", 0)
                start_day = 0 if not start_day else start_day
                start_month = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("month", 0)
                start_month = 0 if not start_month else start_month
                start_year = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("year", 0)
                start_year = 0 if not start_year else start_year
                end_day = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("day", 0)
                end_day = 0 if not end_day else end_day
                end_month = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("month", 0)
                end_month = 0 if not end_month else end_month
                end_year = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("year", 0)
                end_year = 0 if not end_year else end_year

                source = animeData.get("source", "N/A")
                source = 'N/A' if not source else source
                studio = animeData.get("studios", 'N/A')
                studio = 'N/A' if not studio else studio
                if type(studio) is list and studio:
                    studio = studio[0]['name']
                    studio = studio.replace("'", '')
                genresDictList = animeData.get("genres", [])
                genres = [x.get("name", '') for x in genresDictList]

                # filter out uwu
                if "Hentai" in genres or "Ecchi" in genres:
                    continue

                rating = animeData.get("rating", "N/A")
                rating = 'N/A' if not rating else rating
                description = animeData.get("synopsis", "N/A")
                description = 'N/A' if not description else description
                if description and description != 'N/A':
                    description = description.replace("\'", "\'\'")
                year = currYear
                num_episodes = animeData.get("episodes", 0)
                num_episodes = 0 if not num_episodes else num_episodes
                episode_duration = animeData.get("duration", "N/A")
                episode_duration = 'N/A' if not episode_duration else episode_duration
                airing = animeData.get("airing", False)
                airing = False if not airing else airing
                airing = 'FALSE' if not airing else 'TRUE'
                current_status = animeData.get("status", "N/A")
                current_status = 'N/A' if not current_status else current_status
                next_broadcast = "N/A"
                if currYear == 2021 and title_jp in nextAirDateDict:
                    airDate = nextAirDateDict[title_jp]
                    airDateSplit = airDate.split()
                    airYear = airDateSplit[3]
                    airMonth = 1
                    airDay = airDateSplit[1]
                    today = datetime.datetime.now()

                    if int(today.year) <= int(airYear) and int(today.month) <= int(airMonth) and int(today.day) <= int(airDay):
                        next_broadcast = nextAirDateDict[title_jp]
                broadcast_time = animeData.get("broadcast", "N/A")
                broadcast_time = 'N/A' if not broadcast_time else broadcast_time

                score = animeData.get("score", 0)
                score = 0 if not score else score
                scored_by = animeData.get("scored_by", 0)
                scored_by = 0 if not scored_by else scored_by
                rank = animeData.get("rank", 0)
                rank = 0 if not rank else rank
                popularity = animeData.get("popularity", 0)
                popularity = 0 if not popularity else popularity
                favorites = animeData.get("favorites", 0)
                favorites = 0 if not favorites else favorites
                
                image_url = animeData.get("image_url", "N/A")
                image_url = 'N/A' if not image_url else image_url

                if title_jp and title_jp != 'N/A':
                #     title_jp = title_jp.replace('\"', '')
                    title_jp = title_jp.replace("\'", "\'\'")
                newAnime = {
                    'title': title,
                    'title_jp': title_jp,
                    'start_day': start_day,
                    'start_month': start_month,
                    'start_year': start_year,
                    'end_day': end_day,
                    'end_month': end_month,
                    'end_year': end_year,
                    'source': source,
                    'studio': studio,
                    'genres': genres,
                    'rating': rating,
                    'description': description,
                    'season': season,
                    'year': year,
                    'num_episodes': num_episodes,
                    'episode_duration': episode_duration,
                    'airing': airing,
                    'current_status': current_status,
                    'next_broadcast': next_broadcast,
                    'score': score,
                    'scored_by': scored_by,
                    'rank': rank,
                    'popularity': popularity,
                    'favorites': favorites,
                    'image_url': image_url,
                }
                queryString += "("
                queryString += "\'" + title + "\', " + "\'" + title_jp + "\', " + str(start_day) + ", " + str(start_month) + ", " + str(start_year) + ", " + str(end_day) + ", " + str(end_month) + ", " + str(end_year) + "," +  "\n\t"
                queryString += "\'" + source + "\', " + "\'" + studio + "\', "
                queryString += "\'{"
                if genres: 
                    for genre in genres:
                        queryString += "\"" + genre + "\", "
                    queryString = queryString[:-2]
                queryString += "}\', " + "\'" + rating + "\', " + "\n\t"
                queryString += "\'" + description + "\', " + "\n\t"
                queryString += "\'" + season + "\', " + str(year) + ", " + str(num_episodes) + ", " + "\'" + episode_duration + "\', " + airing + ", " + "\'" + current_status + "\', " + "\n\t"
                queryString += "\'" + broadcast_time + "\', " + "\'" + next_broadcast + "\'," + "\n\t"
                queryString += str(score) + ", " + str(scored_by) + ", " + str(rank) + ", " + str(popularity) + ", " + str(favorites) + "," + "\n\t"
                queryString += "\'" + image_url + "\'"
                queryString += "),"
                queryString += "\n\t"
    #         break
    #     break
    # break

queryString = queryString[:queryString.rfind(',')] + ';' + "\n"

with open('../db/docker_postgres_init.sql', 'w') as init_file:
    init_file.write(queryString)
