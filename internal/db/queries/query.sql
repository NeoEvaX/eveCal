-- name: GetUserCharacters :many
SELECT
  u.userid,
  u.characterownerhash,
  c.characterid,
  c.expiry,
  c.portraiturl,
  c.name
FROM
  users u
  INNER JOIN characters c ON u.characterownerhash = c.charaterownerhash
WHERE
  u.characterownerhash = $1;

-- name: CreateUser :exec
INSERT INTO
  users (characterOwnerHash)
VALUES
  ($1);

-- name: GetUser :one
SELECT
  u.userid,
  u.characterownerhash
FROM
  users u
WHERE
  u.characterownerhash = $1
