import datetime
import json
import glob
import urllib3
import xmltodict

# next air date xml file
xml_url = "http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&Fields=TID,Title"
http = urllib3.PoolManager()
response = http.request('GET', xml_url)
data = xmltodict.parse(response.data)
tidDict = {}
for item in data['TitleLookupResponse']['TitleItems']['TitleItem']:
    tidDict[item['Title']] = item['TID']

seasons = ['spring', 'summer', 'fall', 'winter']
current_year = 2021
current_season = 1

seasonDict = {
    'spring': 1,
    'summer': 2,
    'fall': 3,
    'winter': 4
}

animeInsert = 'INSERT INTO anime VALUES \n\t'

airInsert = 'INSERT INTO air_date VALUES \n\t'

for currYear in range(2010, 2022):
    for season in seasons:
        animeFiles = glob.glob('../db/data/'+str(currYear)+str(season)+'/*.json')
        for animeFile in animeFiles:

            with open(animeFile, 'r') as jsonfile:
                animeData = json.load(jsonfile)

                title = animeData.get("title", 'N/A')
                title = 'N/A' if not title else title
                if title and title != 'N/A':
                    title = title.replace("\'", "\'\'")
                title_jp = animeData.get("title_japanese", 'N/A')
                title_jp = 'N/A' if not title_jp else title_jp
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
                
                if title_jp in tidDict:
                    tid = tidDict[title_jp]

                    # AIR TIMES
                    if currYear >= current_year and seasonDict[season] >= current_season:
                        airInsert += "(" + str(tid) + ", \'{"
                        progInfoXml = "http://cal.syoboi.jp/db.php?Command=ProgLookup&TID=" + tid
                        http = urllib3.PoolManager()
                        progInfoResponse = http.request('GET', progInfoXml)
                        progInfoData = xmltodict.parse(progInfoResponse.data)
                        try: 
                            for item in progInfoData['ProgLookupResponse']['ProgItems']['ProgItem']:
                                date_time_now = datetime.datetime.now()
                                date_time_obj = datetime.datetime.strptime(item['StTime'], '%Y-%m-%d %H:%M:%S')
                                if date_time_obj > date_time_now:
                                    airInsert += "\"" + item['StTime'] + "\", "
                        except:
                            print(tid)
                        airInsert += "}"
                        airInsert = airInsert[:airInsert.rfind(', }')] + '}'
                        airInsert += "\'),\n\t"

                else:
                    tid = 0
                    
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
                    title_jp = title_jp.replace("\'", "\'\'")


                animeInsert += "("
                animeInsert += "\'" + title + "\', " + "\'" + title_jp + "\', " + str(tid) + ", " + str(start_day) + ", " + str(start_month) + ", " + str(start_year) + ", " + str(end_day) + ", " + str(end_month) + ", " + str(end_year) + "," +  "\n\t"
                animeInsert += "\'" + source + "\', " + "\'" + studio + "\', "
                animeInsert += "\'{"
                if genres: 
                    for genre in genres:
                        animeInsert += "\"" + genre + "\", "
                    animeInsert = animeInsert[:-2]
                animeInsert += "}\', " + "\'" + rating + "\', " + "\n\t"
                animeInsert += "\'" + description + "\', " + "\n\t"
                animeInsert += "\'" + season + "\', " + str(year) + ", " + str(num_episodes) + ", " + "\'" + episode_duration + "\', " + airing + ", " + "\'" + current_status + "\', " + "\n\t"
                # animeInsert += "\'" + broadcast_time + "\', " + "\'" + next_broadcast + "\'," + "\n\t"
                animeInsert += str(score) + ", " + str(scored_by) + ", " + str(rank) + ", " + str(popularity) + ", " + str(favorites) + "," + "\n\t"
                animeInsert += "\'" + image_url + "\'"
                animeInsert += "),"
                animeInsert += "\n\t"


animeInsert = animeInsert[:animeInsert.rfind(',')] + ';' + "\n"
airInsert = airInsert[:airInsert.rfind(',')] + ';' + "\n"

with open('../db/migration/000002_add_anime.up.sql', 'w') as init_file:
    init_file.write(animeInsert)

with open('../db/migration/000003_add_broadcast_times.up.sql', 'w') as init_file:
    init_file.write(airInsert)
