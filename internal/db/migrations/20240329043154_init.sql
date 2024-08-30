-- +goose Up
CREATE TABLE users (
  userId SERIAL PRIMARY KEY,
  characterOwnerHash text NOT NULL
);

CREATE TABLE characters (
  characterId SERIAL PRIMARY KEY,
  characterOwnerHash text NOT NULL,
  expiry TIMESTAMPTZ NOT NULL,
  portraitUrl text NOT NULL,
  name text NOT NULL
);

CREATE TABLE tokens (
  tokenId SERIAL PRIMARY KEY,
  characterId integer NOT NULL,
  access_token text NOT NULL,
  token_type text NOT NULL,
  refresh_token text NOT NULL,
  expiry TIMESTAMPTZ NOT NULL
);

CREATE TABLE scopes (
  scopeId SERIAL PRIMARY KEY,
  data BYTEA NOT NULL,
  expiry TIMESTAMPTZ NOT NULL
);

CREATE TABLE events (
  eventId SERIAL PRIMARY KEY,
  characterId integer NOT NULL,
  date DATE NOT NULL,
  duration text NOT NULL,
  ownerName text NOT NULL,
  ownerType text NOT NULL,
  response text NOT NULL,
  text text NOT NULL,
  title text NOT NULL
);

-- +goose Down
DROP TABLE events;

DROP TABLE scopes;

DROP TABLE tokens;

DROP TABLE characters;

DROP TABLE users;
