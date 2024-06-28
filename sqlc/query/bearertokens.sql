-- name: CreateBearerToken :exec
INSERT INTO BearerTokens (
    tokenString, validTill, userName
) VALUES (
    $1, $2, $3
);

-- name: RetrieveBearerToken :one
SELECT * FROM BearerTokens
WHERE tokenString = $1 AND valid = True;

-- name: DeleteBearerToken :exec
UPDATE BearerTokens
SET valid = False
WHERE userName = $1 AND valid = True;

-- name: UpdateBearerTokenExpiration :exec
UPDATE BearerTokens
SET validTill = $2
WHERE tokenString = $1 AND valid = True;