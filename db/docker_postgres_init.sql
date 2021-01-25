CREATE TABLE "anime" (
    "title" varchar NOT NULL,
    "title_jp" varchar NOT NULL,
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
    "season" varchar NOT NULL,
    "year" varchar NOT NULL,
    "num_episodes" bigint NOT NULL,
    "episode_duration" varchar NOT NULL,
    "airing" boolean NOT NULL,
    "current_status" varchar NOT NULL,
    "next_broadcast" varchar NOT NULL,
    "score" float NOT NULL,
    "scored_by" bigint NOT NULL,
    "rank" bigint NOT NULL,
    "popularity" bigint NOT NULL,
    "favorites" bigint NOT NULL,
    "image_url" varchar NOT NULL,
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

with anime_json (doc) as (
   values 
    (''::json)
)
insert into anime (
    title,
    title_jp,
    start_day,
    start_month,
    start_year,
    end_day,
    end_month,
    end_year,
    source,
    studio,
    genres,
    rating,
    description,
    season,
    year,
    num_episodes,
    episode_duration,
    airing,
    current_status,
    next_broadcast,
    score,
    scored_by,
    rank,
    popularity,
    favorites,
    image_url,
)
select p.*
from anime_json l
  cross join lateral json_populate_recordset(null::anime, doc) as p
