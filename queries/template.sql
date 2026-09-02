-- name: DbGetTemplate :one
SELECT *
FROM templates
WHERE channel=$1 AND name=$2;