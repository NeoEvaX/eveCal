-- name: GetUserCharacters :many
select u.userid, u.characterownerhash, c.characterid, c.expiry, c.portraiturl, c.name
from users u
inner join characters c on u.characterownerhash = c.charaterownerhash
where u.characterownerhash = $1
limit 1
;

