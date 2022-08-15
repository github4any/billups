-- name: GetScoreboard :many
SELECT * FROM scoreboard
ORDER BY created_at DESC LIMIT 10;

-- name: CreateScoreboard :exec
INSERT INTO scoreboard(results, player, computer, created_at)
VALUES ($1, $2, $3, now());

-- name: DeleteScoreboard :exec
DELETE FROM scoreboard;