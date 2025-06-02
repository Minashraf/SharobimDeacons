CREATE TABLE deacons.attendances
(
    deacon_id BIGINT
        CONSTRAINT attendance_deacons_id_fk
            REFERENCES deacons.deacons
        NOT NULL ,
    event_skill_liturgy_id  INTEGER
        CONSTRAINT attendance_liturgy_event_skill_liturgy_id_fk
            REFERENCES deacons.event_skill_liturgy
        NOT NULL ,
    date      DATE NOT NULL,
    CONSTRAINT attendance_pk
        UNIQUE (deacon_id, date, event_skill_liturgy_id)
);

