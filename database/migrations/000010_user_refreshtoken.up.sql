create table deacons.deacons.user_token
(
    user_id       bigint
        constraint user_token_pk
            primary key,
    refresh_token text      not null,
    updated_at    timestamp not null
);

