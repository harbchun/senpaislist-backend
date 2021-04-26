-- name: CreateAnime :one
INSERT INTO anime (
    title,
    title_jp,
    tid,
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
    image_url
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
    $11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
    $21,$22,$23,$24,$25,$26,$27
) RETURNING *;