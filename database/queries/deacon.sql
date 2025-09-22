-- name: GetDeaconById :one
SELECT
    d.first_name,
    d.last_name,
    d.date_of_birth,
    d.address,
    d.email,
    d.phone_number,
    d.country,
    dr.rank_name
FROM deacons.deacons.deacons d
Join deacons.deacons.deacon_ranks dr on d.deacon_rank_id=dr.id
WHERE d.id = $1;

-- name: GetDeaconSkillById :many
SELECT
    sk.id, sk.skill, ds.score
FROM deacons.deacons.deacon_skill ds
         Join deacons.deacons.skills sk on ds.skill_id = sk.id
WHERE ds.deacon_id = $1;

-- name: GetHistoryServiceByDeaconId :many
SELECT
    skl.skill,
    evnt.event_name,
    lit.liturgy_name,
    att.date
FROM deacons.deacons.attendances att
         Join deacons.deacons.event_skill_liturgy esl on att.event_skill_liturgy_id = esl.id
         Join deacons.deacons.liturgies lit on esl.liturgy_id = lit.id
         Join deacons.deacons.events evnt on esl.event_id = evnt.id
         Join deacons.deacons.skills skl on esl.skill_id = skl.id
WHERE att.deacon_id = $1
OFFSET $2
LIMIT $3;

-- name: CreateDeacon :one
INSERT INTO deacons.deacons.deacons (first_name, last_name, address, email, phone_number, date_of_birth, country, deacon_rank_id)
VALUES (@first_name, @last_name, @address, @email, @phone_number, @date_of_birth, @country, @deacon_rank_id)
    RETURNING id;

-- name: InsertDeaconSkill :exec
INSERT INTO deacons.deacons.deacon_skill (deacon_id, skill_id,score)
VALUES (@deacon_id, @skill_id, @score);

-- name: UpdateDeacon :exec
UPDATE deacons.deacons.deacons SET first_name = $1, last_name = $2, address = $3, email = $4, phone_number = $5, date_of_birth = $6, country = $7, deacon_rank_id = $8
WHERE id=$9;

-- name: DeleteDeacon :exec
DELETE FROM deacons.deacons.deacons where id=$1;

-- name: DeleteDeaconSkill :exec
DELETE From deacons.deacons.deacon_skill where deacon_id=$1;
