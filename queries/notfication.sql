-- name: DbCreateNotification :one
INSERT INTO notifications(id,channel,recipient,template,variables,priority,status)
VALUES ($1,$2,$3,$4,$5,$6,$7)
    RETURNING *;

-- name: DbGetNotificationById :one
SELECT *
FROM notifications
WHERE id=$1;

-- name: DBUpdateNotificationStatus :execresult
UPDATE notifications
SET status=$2
WHERE id=$1
RETURNING *
;