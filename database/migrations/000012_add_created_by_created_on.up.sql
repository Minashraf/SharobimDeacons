ALTER TABLE deacons.attendances
    ADD COLUMN created_on TIMESTAMPTZ,
    ADD COLUMN created_by BIGINT;

UPDATE deacons.attendances
SET created_on = date::timestamp AT TIME ZONE 'Africa/Cairo',
    created_by = 4
WHERE created_on IS NULL;

ALTER TABLE deacons.attendances
    ALTER COLUMN created_on SET DEFAULT now();
