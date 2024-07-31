-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
  userId		integer PRIMARY KEY,
  charaterOwnerHash     text      NOT NULL,
);

CREATE TABLE character (
	characterId		integer PRIMARY KEY,
	charaterOwnerHash	text NOT NULL,
	expiry			TIMESTAMPTZ NOT NULL,
	portraitUrl		text NOT NULL,
	name			text NOT NULL,
);

CREATE TABLE token (
	tokenId 		integer PRIMARY KEY,
	characterId		integer NOT NULL,
	access_token		text NOT NULL,
	token_type 		text NOT NULL,
	refresh_token		text NOT NULL,
	expiry 			TIMESTAMPTZ NOT NULL
);

CREATE TABLE scope (
	scopeId integer PRIMARY KEY,
	data    BYTEA NOT NULL,
	expiry 	TIMESTAMPTZ NOT NULL
);

CREATE TABLE event (
	eventId		integer PRIMARY KEY,
	characterId	integer NOT NULL,
	date		DATE NOT NULL,
	duration	text NOT NULL,
	ownerName	text NOT NULL,
	ownerType	text NOT NULL,
	response	text NOT NULL,
	text		text NOT NULL,
	title		text NOT NULL,
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE event;

DROP TABLE scope;

DROP TABLE token;

DROP TABLE character;

DROP TABLE user;
-- +goose StatementEnd


