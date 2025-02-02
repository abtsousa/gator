-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.*, u.name AS "user"
FROM feeds f JOIN users u ON f.user_id = u.id;

-- name: GetFeed :one
SELECT * FROM feeds WHERE (feeds.url = $1);

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = (SELECT CURRENT_TIMESTAMP), last_fetched_at = (SELECT CURRENT_TIMESTAMP)
WHERE feeds.id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
