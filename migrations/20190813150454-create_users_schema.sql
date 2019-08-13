
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS users;

-- +migrate Down

DROP SCHEMA IF EXISTS users;
