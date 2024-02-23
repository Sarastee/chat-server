-- +goose Up
-- +goose StatementBegin
create table if not exists chats (
    id serial primary key,
    users text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users
-- +goose StatementEnd
