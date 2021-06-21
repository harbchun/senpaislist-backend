CREATE TABLE IF NOT EXISTS broadcast_times(
   syoboi_tid INT,
   time REAL,
   CONSTRAINT fk_syoboi_id FOREIGN KEY (syoboi_tid) REFERENCES airing_informations (syoboi_tid)
);