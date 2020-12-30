CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "user_name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookmarks" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "anime_id" bigint,
  "type" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime" (
  "id" bigserial PRIMARY KEY,
  "title" text NOT NULL,
  "contentType" text NOT NULL,
  "beginDate" text NOT NULL,
  "premieredSeason" text NOT NULL,
  "airing" boolean NOT NULL,
  "currentStatus" text NOT NULL,
  "numEpisodes" integer NOT NULL,
  "episodeDuration" text NOT NULL,
  "broadcastTime" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "users" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("anime_id") REFERENCES "anime" ("id");

-- CREATE INDEX ON "users" ("user_name");

-- CREATE INDEX ON "bookmarks" ("account_id");
