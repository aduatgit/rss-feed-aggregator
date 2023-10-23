-- name: FollowFeed :one
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFollow :exec
DELETE FROM feed_follows 
WHERE id = $1 AND user_id = $2;

-- name: GetUserFeedFollows :many
select * from feed_follows 
WHERE feed_follows.user_id = $1;