CREATE TABLE deacons.liturgies
(
    id           INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT liturgy_pk
            PRIMARY KEY,
    liturgy_name VARCHAR(50) NOT NULL CONSTRAINT  liturgy_pk2 UNIQUE

);

CREATE TABLE deacons.events
(
    id             INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT event_pk
            PRIMARY KEY,
    event_name     VARCHAR(50)
        CONSTRAINT event_pk2
            UNIQUE NOT NULL,
    weight        DECIMAL NOT NULL
);

CREATE TABLE deacons.event_skill_liturgy
(
    id             INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT event_skill_liturgy_pk
            PRIMARY KEY,
    liturgy_id    INTEGER
        CONSTRAINT events_skills_liturgy_id_fk
            REFERENCES deacons.liturgies
        NOT NULL,
    event_id      INTEGER
        CONSTRAINT events_skills_event_id_fk
            REFERENCES deacons.events
        NOT NULL,
    skill_id      INTEGER
        CONSTRAINT events_skills_skills_id_fk
            REFERENCES deacons.skills
        NOT NULL,
    minimum_score DECIMAL NOT NULL,
    CONSTRAINT check_name
        CHECK (event_skill_liturgy.minimum_score BETWEEN 1 AND 10),
    CONSTRAINT event_skill_liturgy_pk_2
        UNIQUE (liturgy_id, event_id, skill_id)
);

