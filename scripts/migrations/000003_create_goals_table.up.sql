create table if not exists goals
(
    id         bigint not null,
    title      text   not null,
    start_at   bigint not null,
    end_at     bigint not null,
    created_at bigint not null,
    constraint goals_pk
        primary key (id)
);
