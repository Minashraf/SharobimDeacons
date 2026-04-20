alter table deacons.deacons
    alter column created_on drop not null,
    alter column created_by drop not null,
    alter column modified_on drop not null,
    alter column modified_by drop not null;

alter table deacons.attendances
    alter column created_on drop not null,
    alter column created_by drop not null;