CREATE TABLE IF NOT EXISTS airing_informations(
   anime_id uuid UNIQUE NOT NULL,
   start_day bigint,
   start_month bigint,
   start_year bigint,
   year INT NOT NULL,
   season VARCHAR NOT NULL,
   num_episodes INT,
   episode_duration VARCHAR,
   airing BOOLEAN,
   syoboi_tid INT,
   CONSTRAINT anime_id_syoboi_id UNIQUE  (anime_id, syoboi_tid),
   CONSTRAINT fk_anime_id_syoboi_id FOREIGN KEY (anime_id, syoboi_tid) REFERENCES animes (id, syoboi_tid),
   CONSTRAINT fk_season FOREIGN KEY (season) REFERENCES seasons (season)
);
 