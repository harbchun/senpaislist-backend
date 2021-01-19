import requests 
import json
import time
import urllib3
import xml.etree.ElementTree as ET

seasons = ['spring', 'summer', 'fall', 'winter']
# xml_url = "http://cal.syoboi.jp/proginfo.xml"
# xml_res = requests.get(xml_url)
# xml_data = ET.parse(xml_res.text)
# print(xml_data)

def main():
    for year in range(2000, 2021):
        for season in seasons:

            with requests.get("https://api.jikan.moe/v3/season/" + str(year) + "/" + season) as seasonResponse:
                time.sleep(4)
                seasonData = seasonResponse.text
                jobj = json.loads(seasonData)

                if not jobj['anime']: continue

                anime_ids = [x['mal_id'] for x in jobj['anime']]
                
                for anime_id in anime_ids: 
                    with requests.get("https://api.jikan.moe/v3/anime/" + str(anime_id)) as animeResponse:
                        time.sleep(4)
                        animeData = animeResponse.text
                        animeObject = json.loads(animeData)

                        print(animeObject["title_japanese"])
                        return

                        animeDict = {}
                        animeDict["title"] = animeObject["title"]
                        animeDict["title_jp"] = animeObject["title_japanese"]
                        animeDict["start_day"] = animeObject["title_japanese"]
                        animeDict["start_month"] = animeObject["title_japanese"]
                        animeDict["start_year"] = animeObject["title_japanese"]
                        animeDict["end_day"] = animeObject["title_japanese"]
                        animeDict["end_month"] = animeObject["title_japanese"]
                        animeDict["end_year"] = animeObject["title_japanese"]
                        animeDict["source"] = animeObject["title_japanese"]
                        animeDict["studio"] = animeObject["title_japanese"]
                        animeDict["genres"] = animeObject["title_japanese"]
                        animeDict["rating"] = animeObject["title_japanese"]
                        animeDict["description"] = animeObject["title_japanese"]
                        animeDict["season"] = animeObject["title_japanese"]
                        animeDict["year"] = animeObject["title_japanese"]
                        animeDict["num_episodes"] = animeObject["title_japanese"]
                        animeDict["episode_duration"] = animeObject["title_japanese"]
                        animeDict["airing"] = animeObject["title_japanese"]
                        animeDict["current_status"] = animeObject["title_japanese"]

                        animeDict["next_broadcast"] = animeObject["title_japanese"]

                        animeDict["score"] = animeObject["score"]
                        animeDict["scored_by"] = animeObject["scored_by"]
                        animeDict["rank"] = animeObject["rank"]
                        animeDict["popularity"] = animeObject["popularity"]
                        animeDict["favorites"] = animeObject["favorites"]
                        animeDict["image_url"] = animeObject["image_url"]

    
if __name__ == "__main__":
    main()
