

CREATE TABLE IF NOT EXISTS tokens(
    id BIGSERIAL PRIMARY KEY,
    token BYTEA NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expiry_at TIMESTAMP NOT NULL,
    scope VARCHAR(30) NOT NULL
);


CREATE INDEX IF NOT EXISTS idx_tokens_token_scope_expiry ON tokens(token, scope, expiry_at);
