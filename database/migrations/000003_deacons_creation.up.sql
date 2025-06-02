CREATE TABLE deacons.deacon_ranks
(
    id        INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT deacon_ranks_pk
            PRIMARY KEY,
    rank_name VARCHAR(50) NOT NULL
        CONSTRAINT deacon_ranks_pk_2
            UNIQUE
);

INSERT INTO deacons.deacon_ranks (rank_name) VALUES ('ابصالتس'),('اغنسطس'),('إبي ذياكون'),('ذياكون'),('أرشي ذياكون');


CREATE TABLE deacons.deacons
(
    id    BIGINT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT deacons_pk
            PRIMARY KEY,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    address     VARCHAR(255),
    email        VARCHAR(255),
    CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    phone_number VARCHAR(20),
    DATE_of_birth DATE,
    country VARCHAR(25) default 'Egypt' NOT NULL ,
    deacon_rank_id INTEGER
        CONSTRAINT deacons_deacon_rank_id_fk
            REFERENCES deacons.deacon_ranks
            NOT NULL,
    CONSTRAINT deacons_pk_2
        UNIQUE (first_name, last_name)
);

CREATE TABLE deacons.skills
(
    id    INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT skills_pk
            PRIMARY KEY,
    skill VARCHAR(50) NOT NULL
        CONSTRAINT skills_pk_2
            UNIQUE
);

CREATE TABLE deacons.deacon_skill
(
    deacon_id  BIGINT
        CONSTRAINT deacon_skill_deacons_id_fk
            REFERENCES deacons.deacons
        NOT NULL,
    skill_id INTEGER
        CONSTRAINT deacon_skill_skills_id_fk
            REFERENCES deacons.skills
        NOT NULL,
    score    decimal NOT NULL ,
    CONSTRAINT check_name
        check (deacon_skill.score between 1 and 10),
    CONSTRAINT deacon_skill_pk
        UNIQUE (deacon_id, skill_id)
);

CREATE TABLE deacons.deacon_ressama
(
    deacon_id      BIGINT
        CONSTRAINT deacon_ressama_deacons_id_fk
            REFERENCES deacons.deacons
        NOT NULL,
    deacon_rank_id INTEGER
        CONSTRAINT deacon_ressama_deacon_rank_id_fk
            REFERENCES deacons.deacon_ranks
        NOT NULL,
    DATE           DATE NOT NULL,
    CONSTRAINT deacon_ressama_pk
        UNIQUE (deacon_id, deacon_rank_id)
);


