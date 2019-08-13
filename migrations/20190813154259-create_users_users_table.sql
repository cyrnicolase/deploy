
-- +migrate Up
CREATE TABLE IF NOT EXISTS users.users (
    id uuid not null primary key,
    username varchar(64) not null unique,
    password varchar(64) not null,
    salt varchar(32) not null,
    email varchar(64) not null,
    phone varchar(18),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);

COMMENT ON COLUMN users.users.username IS '用户名';
COMMENT ON COLUMN users.users.password IS '密码';
COMMENT ON COLUMN users.users.salt IS '盐';
COMMENT ON COLUMN users.users.email IS '邮箱';
COMMENT ON COLUMN users.users.phone IS '手机号';

-- +migrate Down

DROP TABLE users.users;