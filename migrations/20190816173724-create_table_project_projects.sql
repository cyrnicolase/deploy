
-- +migrate Up
CREATE TABLE IF NOT EXISTS project.projects (
    id uuid not null primary key,
    name varchar(64) not null unique,
    alias varchar(32) not null,
    repo varchar(128) not null,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);

COMMENT ON COLUMN project.projects.name IS '项目名字';
COMMENT ON COLUMN project.projects.alias IS '项目别名；服务器可访问路径';
COMMENT ON COLUMN project.projects.repo IS '仓库地址';
-- +migrate Down

DROP TABLE project.projects;