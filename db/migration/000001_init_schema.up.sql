-- CREATE TABLE "users" (
--   "id" bigserial PRIMARY KEY,
--   "user_name" text NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE "bookmarks" (
--   "id" bigserial PRIMARY KEY,
--   "account_id" bigint,
--   "anime_id" bigint,
--   "type" text NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

CREATE TABLE "anime" (
    "title" TEXT,
    "title_jp" TEXT,
    "type" text,
    "source" TEXT, 
    "begin_date" text,
    "end_date" text,
    "genre" text[],
    "season" text,
    "year" INT,
    "airing" boolean,
    "current_status" text,
    "num_episodes" integer,
    "episode_duration" text,
    "broadcast_time" text,
    "next_broadcast" text,
    "score" FLOAT,
    "scored_by" INT,
    "rank" INT,
    "popularity" INT,
    "favorites" INT,
    "image_url" TEXT, 
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

-- ALTER TABLE "bookmarks" ADD FOREIGN KEY ("anime_id") REFERENCES "anime" ("id");

-- CREATE INDEX ON "users" ("user_name");

-- CREATE INDEX ON "bookmarks" ("account_id");
