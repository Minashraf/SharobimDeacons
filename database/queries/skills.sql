-- name: GetSkills :many
SELECT
    *
FROM deacons.deacons.skills;

-- name: GetDependantSkills :many
SELECT DISTINCT
    esl.id as "esl_id",sk.id as "skill_id", sk.skill
FROM deacons.deacons.skills sk join deacons.deacons.event_skill_liturgy esl on sk.id = esl.skill_id
Where esl.liturgy_id = $1 AND esl.event_id=$2;