
CREATE TABLE IF NOT EXISTS messages(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    message TEXT NOT NULL,
    created_at TIME NOT NULL DEFAULT NOW(),
    updated_at TIME NOT NULL DEFAULT NOW(),
    is_enabled BOOLEAN DEFAULT TRUE
);

