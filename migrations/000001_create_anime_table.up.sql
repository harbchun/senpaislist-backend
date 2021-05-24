CREATE TABLE IF NOT EXISTS animes(
   anime_id serial PRIMARY KEY,
   title VARCHAR (50) UNIQUE NOT NULL,
   title_jp VARCHAR (50) NOT NULL,
   anime_description VARCHAR (300) UNIQUE NOT NULL
);
