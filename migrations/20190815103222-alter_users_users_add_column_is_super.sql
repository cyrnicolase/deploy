
-- +migrate Up
ALTER TABLE users.users ADD COLUMN is_super SMALLINT DEFAULT 0;
COMMENT ON COLUMN users.users.is_super IS '是否超级管理员;';
-- +migrate Down

ALTER TABLE users.users DROP COLUMN is_super;
