-- +goose Up
create table chat (
    id serial primary key,
    owner_id int
);

create table chat_user (
    chat_id int not null,
    user_id int not null,
    foreign key (chat_id) references chat (id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table message (
    id bigserial primary key,
    sent_at timestamp not null default now(),
    chat_id int not null,
    user_id int,
    text text not null,
    foreign key (chat_id) references chat (id) ON DELETE CASCADE ON UPDATE CASCADE
);


-- +goose Down
drop table message;
drop table chat_user;
drop table chat;