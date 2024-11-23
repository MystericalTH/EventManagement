-- name: InsertChat :exec
INSERT INTO chatDevAd (adminID, developerID, message, datetime) 
VALUES (?, ?, ?, NOW());

-- name: ListChat :many
SELECT * FROM chatDevAd