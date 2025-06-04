-- name: CreateUser :one
INSERT INTO deacons.deacons.users (email, password)
VALUES (@email, @password)
RETURNING id;

-- name: AssignRole :exec
INSERT INTO deacons.deacons.user_role (user_id, role_id)
VALUES (@user_id, @role_id);

-- name: GetRoles :many
SELECT *
FROM deacons.deacons.roles;

-- name: GetUserByEmail :one
SELECT
    u.id,
    u.email,
    u.password,
    STRING_AGG(r.role, ',') AS roles
FROM deacons.deacons.users u
         JOIN deacons.deacons.user_role ur ON u.id = ur.user_id
         JOIN deacons.deacons.roles r ON r.id = ur.role_id
WHERE u.email = $1
GROUP BY u.id, u.email, u.password;