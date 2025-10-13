-- name: GetEventSkillLiturgy :one
SELECT
    *
FROM deacons.deacons.event_skill_liturgy esl
WHERE esl.id=$1;