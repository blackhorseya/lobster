create table if not exists tasks
(
    id         bigint not null,
    result_id  bigint not null,
    title      text   not null,
    status     int    not null,
    created_at bigint not null,
    constraint tasks_pk
        primary key (id)
);
