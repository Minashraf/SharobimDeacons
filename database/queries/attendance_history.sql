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
ORDER BY att.date DESC, esl.id DESC
OFFSET $2
    LIMIT $3;

-- name: AddHistoryService :exec
INSERT INTO deacons.deacons.attendances (deacon_id, event_skill_liturgy_id, date, created_by) Values($1,$2,$3, $4);

-- name: DeleteHistoryService :exec
DELETE from deacons.deacons.attendances WHERE deacon_id= $1 and event_skill_liturgy_id= $2 and date=$3;

-- name: GetAllHistoryService :many
SELECT
    esl.id as esl_id,
    skl.skill,
    evnt.event_name,
    lit.liturgy_name,
    att.date,
    deacon.first_name,
    deacon.last_name,
    deacon.id as deacon_id
FROM deacons.deacons.deacons deacon
        Join deacons.deacons.attendances att on deacon.id = att.deacon_id
         Join deacons.deacons.event_skill_liturgy esl on att.event_skill_liturgy_id = esl.id
         Join deacons.deacons.liturgies lit on esl.liturgy_id = lit.id
         Join deacons.deacons.events evnt on esl.event_id = evnt.id
         Join deacons.deacons.skills skl on esl.skill_id = skl.id
ORDER BY att.date DESC, esl_id DESC, deacon.id DESC
OFFSET $1
    LIMIT $2;

-- name: GetSuggestion :many
SELECT
    d.id,
    d.first_name,
    d.last_name,
    COALESCE(
            SUM(
                    e.weight
                        * POWER(
                            0.5,
                            EXTRACT(YEAR FROM CURRENT_DATE)
                                - a.year
                          )
            ),
            0
    )::DOUBLE PRECISION AS score
FROM deacons.deacons.deacons d
         JOIN deacons.deacons.deacon_skill ds
              ON d.id = ds.deacon_id
         JOIN deacons.deacons.event_skill_liturgy esl_filter
              ON esl_filter.id = sqlc.arg(eslId)
         LEFT JOIN deacons.deacons.attendances a
                   ON d.id = a.deacon_id
                       AND a.date >= date_trunc('year', CURRENT_DATE) - INTERVAL '4 years'
         LEFT JOIN deacons.deacons.event_skill_liturgy esl
                   ON a.event_skill_liturgy_id = esl.id
         LEFT JOIN deacons.deacons.events e
                   ON esl.event_id = e.id
WHERE ds.skill_id = esl_filter.skill_id
  AND ds.score >= esl_filter.minimum_score
GROUP BY d.id, d.full_name
ORDER BY score, d.full_name
OFFSET sqlc.arg(p_offset)
    LIMIT sqlc.arg(p_limit);