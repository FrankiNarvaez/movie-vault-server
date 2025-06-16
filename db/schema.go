package database

const SCHEMA = `
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	username VARCHAR(25) UNIQUE,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255),
	remember_at TIMESTAMP,
	created_at TIMESTAMP
);

CREATE TABLE favorites (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	imdb_id VARCHAR NOT NULL,
	type VARCHAR(10) NOT NULL,
	created_at TIMESTAMP
);

CREATE TABLE watchlists (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	imdb_id VARCHAR NOT NULL,
	type VARCHAR(10) NOT NULL,
	created_at TIMESTAMP
);
	`

const DROP_SCHEMA = `
	DROP TABLE favorites;
	DROP TABLE watchlists;
	DROP TABLE users;
	`
