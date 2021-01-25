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

allAnime = []
for currYear in range(2010, 2022):
    for season in seasons:
        animeFiles = glob.glob('../db/data/'+str(currYear)+str(season)+'/*.json')
        for animeFile in animeFiles:

            with open(animeFile, 'r') as jsonfile:
                animeData = json.load(jsonfile)

                title = animeData.get("title", 'N/A')
                title = 'N/A' if not title else title
                if title and title != 'N/A':
                    title = title.replace('\"', '')
                    title = title.replace("'", '')
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
                    description = description.replace('\"', '')
                    description = description.replace("'", '')
                year = str(currYear)
                num_episodes = animeData.get("episodes", 0)
                num_episodes = 0 if not num_episodes else num_episodes
                episode_duration = animeData.get("duration", "N/A")
                episode_duration = 'N/A' if not episode_duration else episode_duration
                airing = animeData.get("airing", False)
                airing = False if not airing else airing
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
                    # next_broadcast = nextAirDateDict[title_jp]

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
                    title_jp = title_jp.replace('\"', '')
                    title_jp = title_jp.replace("'", '')
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
                allAnime.append(newAnime)
    
# create the directory to store seasons
# if not os.path.exists('../db/data/'+str(currYear)): os.mkdir('../db/data/'+str(currYear))
with open('../db/data/seed.json', 'w') as outfile:
    json.dump(allAnime, outfile, ensure_ascii=False)
