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