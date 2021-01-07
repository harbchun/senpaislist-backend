-- CREATE TABLE "users" (
--   "id" bigserial PRIMARY KEY,
--   "user_name" VARCHAR NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE "bookmarks" (
--   "id" bigserial PRIMARY KEY,
--   "account_id" bigint,
--   "anime_id" bigint,
--   "type" VARCHAR NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

CREATE TABLE "anime" (
    "title" VARCHAR NOT NULL,
    "title_jp" VARCHAR NOT NULL,
    "show_type" VARCHAR NOT NULL,
    "source" VARCHAR NOT NULL, 
    "begin_date" VARCHAR NOT NULL,
    "end_date" VARCHAR,
    "genre" VARCHAR[],
    "season" VARCHAR NOT NULL,
    "year" BIGINT NOT NULL,
    "airing" BOOLEAN NOT NULL,
    "current_status" VARCHAR NOT NULL,
    "num_episodes" BIGINT NOT NULL,
    "episode_duration" VARCHAR NOT NULL,
    "broadcast_time" VARCHAR NOT NULL,
    "next_broadcast" VARCHAR,
    "score" FLOAT NOT NULL,
    "scored_by" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "popularity" BIGINT NOT NULL,
    "favorites" BIGINT NOT NULL,
    "image_url" VARCHAR NOT NULL, 
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

-- ALTER TABLE "bookmarks" ADD FOREIGN KEY ("anime_id") REFERENCES "anime" ("id");

-- CREATE INDEX ON "users" ("user_name");

-- CREATE INDEX ON "bookmarks" ("account_id");
