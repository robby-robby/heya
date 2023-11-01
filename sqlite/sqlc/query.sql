-- name: GetSettings :one
SELECT *
FROM Settings
LIMIT 1;
-- name: CreateSettings :exec
INSERT INTO Settings (codify, model, editor, temp)
VALUES (?, ?, ?, ?);
-- name: GetConvo :one
SELECT *
FROM Convo
WHERE id = ?;
-- name: ListConvos :many
SELECT *
FROM Convo;
-- name: CreateConvo :exec
INSERT INTO Convo (title, slug, system)
VALUES (?, ?, ?);
-- name: GetMessage :one
SELECT *
FROM Messages
WHERE id = ?;
-- name: ListMessagesByConvo :many
SELECT *
FROM Messages
WHERE convo_id = ?;
-- name: CreateMessage :exec
INSERT INTO Messages (role, msg, convo_id)
VALUES (?, ?, ?);
-- name: GetPin :one
SELECT *
FROM Pins
WHERE id = ?;
-- name: ListPinsByConvo :many
SELECT *
FROM Pins
WHERE convo_id = ?;
-- name: CreatePin :exec
INSERT INTO Pins (convo_id)
VALUES (?);