create table users
(
    id       serial       not null unique,
    nickname varchar(255) not null unique
);

create table chats
(
    id        serial       not null unique,
    chat_name varchar(255) not null unique,
    author    varchar(255) not null
);

create table messages
(
    id       serial                                      not null unique,
    chat_id  int references chats (id) on delete cascade not null,
    nickname varchar(255)                                not null,
    msg      varchar(255)                                not null
);

