create table if not exists results
(
    id         bigint not null,
    goal_id    bigint not null,
    title      text   not null,
    target     int    not null,
    actual     int    not null,
    created_at bigint not null,
    constraint results_pk
        primary key (id)
);
