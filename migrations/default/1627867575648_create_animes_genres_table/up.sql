CREATE TABLE IF NOT EXISTS animes_genres(
   anime_id uuid NOT NULL,
   genre VARCHAR NOT NULL,
   PRIMARY KEY (anime_id, genre),
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id) ON DELETE CASCADE
);
