CREATE TABLE IF NOT EXISTS animes_studios(
   anime_id uuid NOT NULL,
   studio VARCHAR NOT NULL,
   PRIMARY KEY (anime_id, studio),
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id) ON DELETE CASCADE
);
