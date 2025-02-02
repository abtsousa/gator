-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING *
)
SELECT iff.*, f.name AS feed_name, u.name AS user_name
FROM insert_feed_follow iff JOIN feeds f ON iff.feed_id = f.id
JOIN users u ON iff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT f.name, u.name
FROM feed_follows ff JOIN users u ON ff.user_id = u.id
JOIN feeds f on ff.feed_id = f.id
WHERE u.id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
