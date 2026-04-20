alter table deacons.deacons
    alter column created_on set not null,
    alter column created_by set not null,
    alter column modified_on set not null,
    alter column modified_by set not null;

alter table deacons.attendances
    alter column created_on set not null,
    alter column created_by set not null;



