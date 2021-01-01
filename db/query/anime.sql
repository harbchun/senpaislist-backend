-- name: CreateAnime :one
INSERT INTO anime (
    title,
    type,
    source, 
    begin_date,
    end_date,
    premiered_season,
    airing,
    current_status,
    num_episodes,
    episode_duration,
    broadcast_time,
    score,
    scored_by,
    rank,
    popularity,
    favorites
) VALUES (
    $1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16
);