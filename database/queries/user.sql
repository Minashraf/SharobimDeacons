-- name: CreateUser :one
INSERT INTO deacons.deacons.users (email, password)
VALUES (@email, @password)
RETURNING id;

-- name: AssignRole :exec
INSERT INTO deacons.deacons.user_role (user_id, role_id)
VALUES (@user_id, @role_id);

-- name: GetRoles :many
SELECT * FROM deacons.deacons.roles;