# 发布平台
## 数据库有如下一些表

### 用户相关
users.users
- id
- username
- password
- salt
- email
- phone

users.user_privileges
- id
- user_id
- privilege

### 项目相关
project.projects
- id
- name
- alias
- repo

project.servers
- id
- name
- host
- port
- user
- password
- is_ssl

project.project_servers
- id
- project_id
- server_id


