
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE devices (
id int,
devise_token varchar(320),
user_id int,
CONSTRAINT devices_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
PRIMARY KEY(devise_token));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE devices;
