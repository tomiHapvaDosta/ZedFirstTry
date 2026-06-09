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
