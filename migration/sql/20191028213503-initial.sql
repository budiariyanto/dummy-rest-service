-- +migrate Up
CREATE TABLE todo_list (id BIGINT primary key auto_increment, name varchar(100), status TINYINT(1));
-- +migrate Down
DROP TABLE todo_list;
