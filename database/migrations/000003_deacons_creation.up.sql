create table deacons.deacons
(
    id    bigint generated always as identity
        constraint deacons_pk
            primary key,
    first_name  varchar(255) not null,
    last_name   varchar(255) not null,
    address     varchar(255),
    email        VARCHAR(255),
    CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    phone_number varchar(20),
    date_of_birth date,
    country varchar(25) default 'Egypt'
);

create table deacons.skills
(
    id    integer generated always as identity
        constraint skills_pk
            primary key,
    skill varchar(50) not null
        constraint skills_pk_2
            unique
);

create table deacons.deacon_skill
(
    deacon_id  bigint
        constraint deacon_skill_deacons_id_fk
            references deacons.deacons,
    skill_id integer
        constraint deacon_skill_skills_id_fk
            references deacons.skills,
    score    decimal,
    constraint check_name
        check (deacon_skill.score between 1 and 10)
);


