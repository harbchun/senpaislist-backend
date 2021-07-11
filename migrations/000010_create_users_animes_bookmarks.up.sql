CREATE TABLE IF NOT EXISTS users_animes_bookmarks(
   user_id uuid NOT NULL,
   anime_id uuid NOT NULL,
   bookmark_type VARCHAR NOT NULL,
   PRIMARY KEY (user_id, anime_id),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
   CONSTRAINT fk_anime_id FOREIGN KEY (anime_id) REFERENCES animes (id) ON DELETE CASCADE
);
