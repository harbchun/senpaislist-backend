from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import Column, String, Integer, Boolean, Float
from sqlalchemy.dialects import postgresql

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

# connection to the db
engine = create_engine('postgres://postgres:championsclub123@localhost:5432/postgres?sslmode=disable')
Session = sessionmaker(bind=engine)
Base = declarative_base()

class Anime(Base):
    __tablename__ = 'anime'
    id = Column(Integer, primary_key=True)
    title = Column(String)
    title_jp = Column(String)
    start_day = Column(Integer)
    start_month = Column(Integer)
    start_year = Column(Integer)
    end_day = Column(Integer)
    end_month = Column(Integer)
    end_year = Column(Integer)
    source = Column(String)
    studio = Column(String)
    genres = Column(postgresql.ARRAY(String))
    rating = Column(String)
    description = Column(String)
    season = Column(String)
    year = Column(String)
    num_episodes = Column(Integer)
    episode_duration = Column(String)
    airing = Column(Boolean)
    current_status = Column(String)
    next_broadcast = Column(String)
    score = Column(Float)
    scored_by = Column(Integer)
    rank = Column(Integer)
    popularity = Column(Integer)
    favorites = Column(Integer)
    image_url = Column(String)
    def __init__(self, title, title_jp, start_day, start_month, start_year,
                end_day, end_month, end_year, source, studio, genres, rating,
                description, season, year, num_episodes, episode_duration,
                airing, current_status, next_broadcast, score, scored_by, 
                rank, popularity, favorites, image_url):
        self.title = title
        self.title_jp = title_jp
        self.start_day = start_day
        self.start_month = start_month
        self.start_year = start_year
        self.end_day = end_day
        self.end_month = end_month
        self.end_year = end_year
        self.source = source
        self.studio = studio
        self.genres = genres
        self.rating = rating
        self.description = description
        self.season = season
        self.year = year
        self.num_episodes = num_episodes
        self.episode_duration = episode_duration
        self.airing = airing
        self.current_status = current_status
        self.next_broadcast = next_broadcast
        self.score = score
        self.scored_by = scored_by
        self.rank = rank
        self.popularity = popularity
        self.favorites = favorites
        self.image_url = image_url


seasons = ['spring', 'summer', 'fall', 'winter']
Base.metadata.create_all(engine)
newSession = Session()
for year in range(2021, 2022):
    for season in seasons:
        animeFiles = glob.glob('../db/data/'+str(year)+str(season)+'/*.json')
        for animeFile in animeFiles:

            with open(animeFile, 'r') as jsonfile:
                animeData = json.load(jsonfile)

                title = animeData.get("title", 'N/A')
                title = 'N/A' if not title else title
                title_japanese = animeData.get("title_japanese", 'N/A')
                title_japanese = 'N/A' if not title_japanese else title_japanese

                startD = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("day", 0)
                startD = 0 if not startD else startD
                startM = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("month", 0)
                startM = 0 if not startM else startM
                startY = animeData.get("aired", {}).get("prop", {}).get("from", {}).get("year", 0)
                startY = 0 if not startY else startY
                endD = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("day", 0)
                endD = 0 if not endD else endD
                endM = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("month", 0)
                endM = 0 if not endM else endM
                endY = animeData.get("aired", {}).get("prop", {}).get("to", {}).get("year", 0)
                endY = 0 if not endY else endY

                source = animeData.get("source", "N/A")
                source = 'N/A' if not source else source
                studio = animeData.get("studios", 'N/A')
                studio = 'N/A' if not studio else studio
                if type(studio) is list and studio:
                    studio = studio[0]['name']
                genresDictList = animeData.get("genres", [])
                genres = [x.get("name", '') for x in genresDictList]

                # filter out uwu
                if "Hentai" in genres or "Ecchi" in genres:
                    continue

                rating = animeData.get("rating", "N/A")
                rating = 'N/A' if not rating else rating
                description = animeData.get("synopsis", "N/A")
                description = 'N/A' if not description else description
                num_episodes = animeData.get("episodes", 0)
                num_episodes = 0 if not num_episodes else num_episodes
                duration = animeData.get("duration", "N/A")
                duration = 'N/A' if not duration else duration
                airing = animeData.get("airing", False)
                airing = False if not airing else airing
                current_status = animeData.get("status", "N/A")
                current_status = 'N/A' if not current_status else current_status
                next_air = "N/A"
                if year == 2021 and title_japanese in nextAirDateDict:
                    airDate = nextAirDateDict[title_japanese]
                    airDateSplit = airDate.split()
                    airYear = airDateSplit[3]
                    airMonth = 1
                    airDay = airDateSplit[1]
                    today = datetime.datetime.now()

                    if int(today.year) <= int(airYear) and int(today.month) <= int(airMonth) and int(today.day) <= int(airDay):
                        next_air = nextAirDateDict[title_japanese]
                    # next_air = nextAirDateDict[title_japanese]

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

                # print(
                #     animeData['title'],'\n',
                #     animeData['title_japanese'],'\n',
                #     animeData["aired"]["prop"]["from"]["day"],'\n',
                #     animeData["aired"]["prop"]["from"]["month"],'\n',
                #     animeData["aired"]["prop"]["from"]["year"],'\n',
                #     endD,'\n',
                #     endM,'\n',
                #     endY,'\n',
                #     animeData["source"],'\n',
                #     studio,'\n',
                #     [x["name"] for x in animeData["genres"]],'\n',
                #     animeData["rating"],'\n',
                #     animeData["synopsis"],'\n',
                #     season,'\n',
                #     str(year),'\n',
                #     animeData["episodes"],'\n',
                #     animeData["duration"],'\n',
                #     animeData["airing"],'\n',
                #     animeData["status"],'\n',
                #     "N/A",'\n',
                #     animeData["score"],'\n',
                #     animeData["scored_by"],'\n',
                #     animeData["rank"],'\n',
                #     animeData["popularity"],'\n',
                #     animeData["favorites"],'\n',
                #     animeData["image_url"],'\n'
                # )

                newAnime = Anime(
                    title,
                    title_japanese,
                    startD,
                    startM,
                    startY,
                    endD,
                    endM,
                    endY,
                    source,
                    studio,
                    genres,
                    rating,
                    description,
                    season,
                    str(year),
                    num_episodes,
                    duration,
                    airing,
                    current_status,
                    "N/A",
                    score,
                    scored_by,
                    rank,
                    popularity,
                    favorites,
                    image_url
                )
                newSession.merge(newAnime)
                
newSession.commit()
newSession.close()