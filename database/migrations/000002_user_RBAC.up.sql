create table deacons.roles
(
    id   int generated always as identity
        constraint roles_pk
            primary key,
    role varchar(50) not null
        constraint roles_pk_2
            unique
);

create table deacons.users
(
    id       bigint generated always as identity
        constraint users_pk
            primary key,
    email    varchar(255)
        constraint users_pk_2
            unique,
    password varchar(255)
);

create table deacons.user_role
(
    user_id bigint
        constraint user_role_users_id_fk
            references deacons.users,
    role_id integer
        constraint user_role_roles_id_fk
            references deacons.roles,
    constraint user_role_pk
        unique (user_id, role_id)
);

INSERT INTO deacons.roles (role) VALUES ('Super admin'),('admin'), ('sharobim servant'), ('deacon')