-- name: GetEvents :many
SELECT DISTINCT
    e.id, event_name
FROM deacons.deacons.events e join deacons.deacons.event_skill_liturgy esl on e.id = esl.event_id
Where esl.liturgy_id = $1;