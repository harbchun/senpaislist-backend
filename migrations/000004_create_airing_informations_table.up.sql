CREATE TABLE IF NOT EXISTS airing_informations(
   anime_id uuid NOT NULL,
   start_day bigint,
   start_month bigint,
   start_year bigint,
   year INT NOT NULL,
   season VARCHAR NOT NULL,
   num_episodes INT,
   episode_duration VARCHAR,
   airing BOOLEAN,
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id),
   CONSTRAINT fk_season FOREIGN KEY (season) REFERENCES seasons (season),
   CONSTRAINT fk_year FOREIGN KEY (year) REFERENCES years (year)
);
