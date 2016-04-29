
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE sessions (
id int,
user_id int,
CONSTRAINT sessions_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
devise_token varchar(320),
CONSTRAINT sessions_devices_key FOREIGN KEY(devise_token)
REFERENCES devices(devise_token));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE sessions;
