ALTER TABLE deacons.deacons
    ADD COLUMN created_on TIMESTAMPTZ,
    ADD COLUMN created_by BIGINT,
    ADD COLUMN modified_on TIMESTAMPTZ,
    ADD COLUMN modified_by BIGINT;

UPDATE deacons.deacons
SET created_on = now() AT TIME ZONE 'Africa/Cairo',
    created_by = 4,
    modified_on = now() AT TIME ZONE 'Africa/Cairo',
    modified_by = 4
WHERE created_on IS NULL;

ALTER TABLE deacons.deacons
    ALTER COLUMN created_on SET DEFAULT now(),
    ALTER COLUMN modified_on SET DEFAULT now();

-- 1. Create the trigger function
CREATE OR REPLACE FUNCTION update_modified_on()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_on = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 2. Attach the trigger to the profiles table
CREATE OR REPLACE TRIGGER set_modified_on
    BEFORE UPDATE ON deacons.deacons
    FOR EACH ROW
EXECUTE FUNCTION update_modified_on();