-- name: CreatePost :one
INSERT INTO posts (id, user_id, title, body, published, created_at, updated_at)
    VALUES (
        gen_random_uuid(),
        $1,
        $2,
        $3,
        $4,
        NOW(),
        NOW()
    )
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts WHERE published = true ORDER BY created_at DESC;

-- name: GetPost :one
SELECT * FROM posts WHERE published = true AND id=$1;

-- name: UpdatePost :one
UPDATE posts
SET title = $1 AND body = $2
WHERE id = $3
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;
