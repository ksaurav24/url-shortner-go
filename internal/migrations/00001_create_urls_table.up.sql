CREATE TYPE url_status AS ENUM ('active', 'disabled');

CREATE TABLE urls (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ
    status url_status NOT NULL DEFAULT 'active'
);