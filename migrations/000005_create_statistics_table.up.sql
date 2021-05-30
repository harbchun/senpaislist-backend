CREATE TABLE IF NOT EXISTS statistics(
   anime_id uuid NOT NULL,
   score int,
   scored_by int,
   rank int,
   popularity INT,
   favorites int,
   rating INT,
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id)
);
