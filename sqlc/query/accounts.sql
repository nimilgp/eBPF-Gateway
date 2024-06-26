-- name: CreateAccount :exec
INSERT INTO Accounts (
    userName, passwdHash, powerLevel, firstName, lastName 
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: RetrieveAccount :one
SELECT * FROM Accounts
WHERE userName = $1 AND valid = True;

-- name: DeleteAccount :exec
UPDATE Accounts
SET valid = False
WHERE userName = $1 AND valid = True;

-- name: UpdateAccountPowerLevel :exec
UPDATE Accounts
SET powerLevel = $2
WHERE userName = $1 AND valid = True;

-- name: UpdatePasswdHash :exec
UPDATE Accounts
SET passwdHash = $2
WHERE userName = $1 AND valid = True;