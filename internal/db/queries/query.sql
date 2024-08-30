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

-- name: GetCharacter :one
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
  c.characterid = $1;

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
  u.characterownerhash = $1;

-- name: GetCharacterTokens :many
SELECT
  t.tokenId,
  t.characterId,
  t.access_token,
  t.token_type,
  t.refresh_token,
  t.expiry
FROM
  tokens t
WHERE
  t.characterId = $1;

-- name: CreateCharacter :exec
INSERT INTO
  characters (characterOwnerHash, expiry, portraitUrl, name)
VALUES
  ($1, $2, $3, $4);

-- name: CreateToken :exec
INSERT INTO
  tokens (
    characterId,
    access_token,
    token_type,
    refresh_token,
    expiry
  )
VALUES
  ($1, $2, $3, $4, $5);

-- name: CreateScope :exec
INSERT INTO
  scopes (data, expiry)
VALUES
  ($1, $2);
