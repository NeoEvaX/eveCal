-- name: GetUserCharacters :many
select u.userid, u.characterownerhash, c.characterid, c.expiry, c.portraiturl, c.name,
from user u
inner join character c on u.characterownerhash = c.charaterownerhash
where characterownerhash = $1
limit 1
;

