import requests 
import json
import time
import os

seasons = ['spring', 'summer', 'fall', 'winter']

def main():
    for year in range(2012, 2014):
        print('year: ', year)
        for season in seasons:
            print('season: ', season)
            with requests.get("https://api.jikan.moe/v3/season/" + str(year) + "/" + season) as seasonResponse:
                time.sleep(4)
                
                # create the directory to store seasons
                if not os.path.exists('../db/data/'+str(year)+str(season)): os.mkdir('../db/data/'+str(year)+str(season))

                seasonData = seasonResponse.text
                jobj = json.loads(seasonData)

                if not jobj['anime']: continue

                anime_ids = [x['mal_id'] for x in jobj['anime']]
                
                for anime_id in anime_ids: 
                    with requests.get("https://api.jikan.moe/v3/anime/" + str(anime_id)) as animeResponse:
                        time.sleep(4)
                        animeJson = animeResponse.json()
                        # animeData = animeResponse.text
                        # animeObject = json.loads(animeData)
                        with open('../db/data/'+str(year)+str(season)+'/'+str(year)+'_'+str(season)+'_'+str(anime_id)+'.json', 'w') as outfile:
                            json.dump(animeJson, outfile)


    
if __name__ == "__main__":
    main()
