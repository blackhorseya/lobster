create table if not exists users
(
    id         bigint not null,
    email      text   not null,
    password   text   not null,
    token      text   not null,
    created_at bigint not null,
    constraint users_pk
        primary key (id)
);
