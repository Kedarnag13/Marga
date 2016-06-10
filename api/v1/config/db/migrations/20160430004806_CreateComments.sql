
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE comments (
id SERIAL,
description varchar(2083),
user_id int,
CONSTRAINT comments_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
issue_id int,
CONSTRAINT comments_issues_key FOREIGN KEY(issue_id)
REFERENCES issues(id),
PRIMARY KEY(id));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE comments;
