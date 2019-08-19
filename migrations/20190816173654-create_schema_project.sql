
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS project;
-- +migrate Down
DROP SCHEMA project;
