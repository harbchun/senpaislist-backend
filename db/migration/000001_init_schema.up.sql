CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "hashed_password" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookmarks" (
    "owner" varchar NOT NULL,
    "anime_id" bigint,
    "bookmark_type" varchar NOT NULL,
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime" (
    "title" varchar NOT NULL,
    "title_jp" varchar NOT NULL,
    "tid" bigint NOT NULL,
    "start_day" bigint NOT NULL,
    "start_month" bigint NOT NULL,
    "start_year" bigint NOT NULL,
    "end_day" bigint NOT NULL,
    "end_month" bigint NOT NULL,
    "end_year" bigint NOT NULL,
    "source" varchar NOT NULL,
    "studio" varchar NOT NULL,
    "genres" varchar[] NOT NULL,
    "rating" varchar NOT NULL,
    "description" varchar NOT NULL,
    "season_id" varchar NOT NULL,
    "year" bigint NOT NULL,
    "num_episodes" bigint NOT NULL,
    "episode_duration" varchar NOT NULL,
    "airing" boolean NOT NULL,
    "current_status" varchar NOT NULL,
    "score" float NOT NULL,
    "scored_by" bigint NOT NULL,
    "rank" bigint NOT NULL,
    "popularity" bigint NOT NULL,
    "favorites" bigint NOT NULL,
    "image_url" varchar NOT NULL,
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "broadcast_times" (
    "tid" bigint NOT NULL,
    "times" varchar[] NOT NULL
);

CREATE TABLE "seasons" (
    "season_id" varchar NOT NULL,
    "season_name" varchar NOT NULL,
    "start_month" smallint NOT NULL,
    "end_month" smallint NOT NULL
);



ALTER TABLE "bookmarks" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("anime_id") REFERENCES "anime" ("id");

CREATE INDEX ON "bookmarks" ("owner");

CREATE INDEX ON "bookmarks" ("anime_id");
