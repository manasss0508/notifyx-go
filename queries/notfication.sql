-- name: DbCreateNotification :one
INSERT INTO notifications(id,channel,recipient,template,name,priority,status)
VALUES ($1,$2,$3,$4,$5,$6,$7)
    RETURNING *;

-- name: DbGetNotificationById :one
SELECT *
FROM notifications
WHERE id=$1;