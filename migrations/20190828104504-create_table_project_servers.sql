
-- +migrate Up
CREATE TABLE IF NOT EXISTS project.servers (
    id uuid not null primary key,
    addr varchar(32) not null,
    port integer not null,
    uname varchar(32) not null,
    passwd varchar(64),
    sshkey smallint not null default 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

COMMENT ON TABLE project.servers IS '服务器';
COMMENT ON COLUMN project.servers.addr IS 'IP地址';
COMMENT ON COLUMN project.servers.port IS '端口';
COMMENT ON COLUMN project.servers.uname IS '登录用户';
COMMENT ON COLUMN project.servers.passwd IS '登录密码';
COMMENT ON COLUMN project.servers.sshkey IS '是否ssh登录秘钥';

-- +migrate Down
DROP TABLE project.servers;