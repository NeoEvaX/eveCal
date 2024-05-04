-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
  userId    integer PRIMARY KEY,
  charaterOwnerHash     text      NOT NULL,
);

CREATE TABLE character (
	characterId          integer PRIMARY KEY,
	charaterOwnerHash    text NOT NULL,
	expiry               TIMESTAMPTZ NOT NULL
);

CREATE TABLE token (
	tokenId 		integer PRIMARY KEY,
  	characterId 	integer NOT NULL,
	access_token  	text NOT NULL,
	token_type 		text NOT NULL,
  	refresh_token 	text NOT NULL,
	expiry 			TIMESTAMPTZ NOT NULL
);

CREATE TABLE scope (
	scopeId integer PRIMARY KEY,
	data    BYTEA NOT NULL,
	expiry 	TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;

DROP TABLE sessions;

DROP INDEX sessions_expiry_idx;
-- +goose StatementEnd
