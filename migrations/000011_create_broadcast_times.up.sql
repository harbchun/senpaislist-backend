CREATE TABLE IF NOT EXISTS broadcast_times(
   anime_id uuid NOT NULL,
   syoboi_tid INT,
   time REAL,
   CONSTRAINT fk_anime_id_syoboi_id FOREIGN KEY (anime_id, syoboi_tid) REFERENCES airing_informations (anime_id, syoboi_tid)
);