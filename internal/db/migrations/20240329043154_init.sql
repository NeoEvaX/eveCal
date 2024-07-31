-- +goose Up
create table users (
	userId			integer	PRIMARY KEY,
	characterOwnerHash	text	NOT NULL
);

create table characters (
	characterId		integer PRIMARY KEY,
	characterOwnerHash	text NOT NULL,
	expiry			TIMESTAMPTZ NOT NULL,
	portraitUrl		text NOT NULL,
	name			text NOT NULL
);

create table tokens (
	tokenId 		integer PRIMARY KEY,
	characterId		integer NOT NULL,
	access_token		text NOT NULL,
	token_type 		text NOT NULL,
	refresh_token		text NOT NULL,
	expiry 			TIMESTAMPTZ NOT NULL
);

create table scopes (
	scopeId integer PRIMARY KEY,
	data    BYTEA NOT NULL,
	expiry 	TIMESTAMPTZ NOT NULL
);

create table events (
	eventId		integer PRIMARY KEY,
	characterId	integer NOT NULL,
	date		DATE NOT NULL,
	duration	text NOT NULL,
	ownerName	text NOT NULL,
	ownerType	text NOT NULL,
	response	text NOT NULL,
	text		text NOT NULL,
	title		text NOT NULL
);
-- +goose Down
drop table events;

drop table scopes;

drop table tokens;

drop table characters;

drop table users;

