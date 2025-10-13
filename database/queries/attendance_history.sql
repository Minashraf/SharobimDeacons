-- name: GetHistoryServiceByDeaconId :many
SELECT
    esl.id,
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
ORDER BY att.date DESC
OFFSET $2
    LIMIT $3;

-- name: AddHistoryService :exec
INSERT INTO deacons.deacons.attendances (deacon_id, event_skill_liturgy_id, date) Values($1,$2,$3);

-- name: DeleteHistoryService :exec
DELETE from deacons.deacons.attendances WHERE deacon_id= $1 and event_skill_liturgy_id= $2 and date=$3;
