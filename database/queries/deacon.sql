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