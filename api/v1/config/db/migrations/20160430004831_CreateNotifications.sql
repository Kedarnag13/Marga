
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notifications (
id SERIAL,
message varchar(2083),
sender_id int,
CONSTRAINT notifications_users_sender_key FOREIGN KEY(sender_id)
REFERENCES users(id),
reciever_id int,
CONSTRAINT notifications_users_reciever_key FOREIGN KEY(reciever_id)
REFERENCES users(id));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notifications;
