-- name: GetEventSkillLiturgy :one
SELECT
    *
FROM deacons.deacons.event_skill_liturgy esl
WHERE esl.id=$1;

-- name: GetRemainingEventSkillLiturgyCapacity :one
SELECT
    esl.capacity - count(deacon_id)
FROM deacons.deacons.event_skill_liturgy esl left join deacons.deacons.attendances att on esl.id = att.event_skill_liturgy_id
    and EXTRACT(YEAR FROM att.date) = EXTRACT(YEAR FROM sqlc.arg(service_date)::date)
WHERE esl.id = sqlc.arg(eslId)
GROUP BY esl.id;