
-- +migrate Up
CREATE TABLE IF NOT EXISTS users.user_privileges (
    id uuid not null primary key,
    user_id uuid not null,
    privilege varchar(128) not null,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    constraint uk_user_id_privilege unique (user_id, privilege)
);
-- +migrate Down
DROP TABLE IF EXISTS users.user_privileges;
