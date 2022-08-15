
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS scoreboard (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    results VARCHAR NOT NULL,
    player INT4 NOT NULL,
    computer INT4 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
CREATE INDEX scoreboard_results ON scoreboard USING BTREE (results);
-- +migrate Down
DROP TABLE IF EXISTS scoreboard;
