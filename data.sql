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
CREATE TABLE devices (
id int,
devise_token varchar(320),
user_id int,
CONSTRAINT devices_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
PRIMARY KEY(devise_token));
CREATE TABLE sessions (
id int,
user_id int,
CONSTRAINT sessions_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
devise_token varchar(320),
CONSTRAINT sessions_devices_key FOREIGN KEY(devise_token)
REFERENCES devices(devise_token));
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
CREATE TABLE notifications (
id SERIAL,
message varchar(2083),
sender_id int,
CONSTRAINT notifications_users_sender_key FOREIGN KEY(sender_id)
REFERENCES users(id),
reciever_id int,
CONSTRAINT notifications_users_reciever_key FOREIGN KEY(reciever_id)
REFERENCES users(id));
