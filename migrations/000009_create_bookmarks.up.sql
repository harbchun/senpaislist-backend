CREATE TABLE IF NOT EXISTS bookmarks(
   id SERIAL PRIMARY KEY,
   type VARCHAR NOT NULL UNIQUE
);
