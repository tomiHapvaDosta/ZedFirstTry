-- +goose Up
CREATE TABLE
    posts (
        id UUID PRIMARY KEY,
        user_id UUID REFERENCES users(id) ON DELETE CASCADE,
        title TEXT NOT NULL,
        body TEXT NOT NULL,
        published BOOL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );

-- +goose Down
DROP TABLE posts;
