create table if not exists goals
(
    id         bigint not null,
    user_id    bigint not null,
    title      text   not null,
    start_at   bigint not null,
    end_at     bigint not null,
    created_at bigint not null,
    constraint goals_pk
        primary key (id)
);

create table if not exists results
(
    id         bigint not null,
    user_id    bigint not null,
    goal_id    bigint not null,
    title      text   not null,
    target     int    not null,
    actual     int    not null,
    created_at bigint not null,
    constraint results_pk
        primary key (id)
);
