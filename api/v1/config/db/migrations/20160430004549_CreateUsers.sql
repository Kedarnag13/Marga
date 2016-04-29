
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
id SERIAL,
name varchar(100),
username varchar(100),
email varchar(320),
mobile_number varchar(100),
latitude float,
longitude float,
password varchar(100),
password_confirmation varchar(100),
city varchar(100) DEFAULT 'MYSURU',
device_token varchar(320),
created_at timestamptz,
my_points int,
type varchar(100),
PRIMARY KEY(id),
UNIQUE (username, mobile_number));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
