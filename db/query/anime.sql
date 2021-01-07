-- name: CreateAnime :one
INSERT INTO anime (
    title,
    title_jp,
    show_type,
    source, 
    begin_date,
    end_date,
    genre,
    season,
    year,
    airing,
    current_status,
    num_episodes,
    episode_duration,
    broadcast_time,
    next_broadcast,
    score,
    scored_by,
    rank,
    popularity,
    favorites,
    image_url
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21
) RETURNING *;