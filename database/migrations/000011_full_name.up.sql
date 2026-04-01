alter table deacons.deacons
    add full_name text generated always as (deacons.first_name || ' ' || deacons.last_name) stored;
