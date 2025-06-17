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
	tmdb_id VARCHAR NOT NULL,
	type VARCHAR(10) NOT NULL,
	created_at TIMESTAMP
);

CREATE TABLE watchlists (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR NOT NULL,
	created_at TIMESTAMP
);

CREATE TABLE item_watchlists (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	watchlist_id UUID REFERENCES watchlists(id) ON DELETE CASCADE,
	tmdb_id VARCHAR NOT NULL,
	type VARCHAR(10) NOT NULL,
	status INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_favorites_user_id ON favorites(user_id);
CREATE INDEX IF NOT EXISTS idx_favorites_tmdb_id ON favorites(tmdb_id);
CREATE INDEX IF NOT EXISTS idx_watchlists_user_id ON watchlists(user_id);
CREATE INDEX IF NOT EXISTS idx_item_watchlists_watchlist_id ON item_watchlists(watchlist_id);
CREATE INDEX IF NOT EXISTS idx_item_watchlists_tmdb_id ON item_watchlists(tmdb_id);
	`
