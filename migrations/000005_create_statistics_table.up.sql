CREATE TABLE IF NOT EXISTS statistics(
   anime_id uuid NOT NULL,
   score int NOT NULL,
   scored_by int NOT NULL,
   rank int NOT NULL,
   popularity INT NOT NULL,
   favorites int NOT NULL,
   rating INT,
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id)
);
