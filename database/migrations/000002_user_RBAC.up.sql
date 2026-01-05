CREATE TABLE deacons.roles
(
    id   INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT roles_pk
            PRIMARY KEY,
    role VARCHAR(50) NOT NULL
        CONSTRAINT roles_pk_2
            UNIQUE
);

CREATE TABLE deacons.users
(
    id       bigint generated always as identity
        CONSTRAINT users_pk
            PRIMARY KEY,
    email    VARCHAR(255)
        CONSTRAINT users_pk_2
            UNIQUE
        NOT NULL ,
    CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE deacons.user_role
(
    user_id bigint
        CONSTRAINT user_role_users_id_fk
            REFERENCES deacons.users
        NOT NULL,
    role_id integer
        CONSTRAINT user_role_roles_id_fk
            REFERENCES deacons.roles
        NOT NULL,
    CONSTRAINT user_role_pk
        UNIQUE (user_id, role_id)
);

INSERT INTO deacons.roles (role) VALUES ('Super admin'),('admin'), ('sharobim servant'), ('deacon')