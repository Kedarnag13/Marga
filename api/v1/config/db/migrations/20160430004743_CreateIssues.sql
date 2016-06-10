
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE issues (
id SERIAL,
name varchar(320),
type varchar(100),
description varchar(2083),
latitude float,
longitude float,
image varchar(2083),
status boolean,
address varchar(2083),
created_at timestamptz,
user_id int,
CONSTRAINT issues_users_key FOREIGN KEY (user_id)
REFERENCES users(id),
PRIMARY KEY(id));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE issues;
