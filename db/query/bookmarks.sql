-- name: CreateBookmark :one
INSERT INTO bookmarks (
    owner,
    anime_id,
    bookmark_type
) VALUES (
    $1,$2,$3
) RETURNING *;

