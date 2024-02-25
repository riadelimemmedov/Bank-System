-- name: CreateAccount :one
INSERT INTO
    accounts (name, owner, balance, currency)
VALUES
    ($1, $2, $3, $4) RETURNING *;


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateAccount :one
UPDATE accounts
    set name = $2,
    balance = $3
WHERE id = $1
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;