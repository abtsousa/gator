-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT p.*
FROM posts p
JOIN feeds f ON p.feed_id = f.id
JOIN users u ON f.user_id = u.id
ORDER BY p.updated_at DESC LIMIT $1;
