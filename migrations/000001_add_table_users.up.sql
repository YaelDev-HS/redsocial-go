CREATE EXTENSION IF NOT EXISTS citext;


CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    email CITEXT NOT NULL UNIQUE,
    password BYTEA NOT NULL,
    nickname TEXT NOT NULL,
    username CITEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    is_enabled BOOLEAN DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS idx_users_nickname ON users(nickname);
