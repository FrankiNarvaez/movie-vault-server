package database

const SCHEMA = `
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	username VARCHAR(25) UNIQUE,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255),
	remember_at TIMESTAMP
);
	`

const DROP_SCHEMA = `
	DROP TABLE users;
	`
