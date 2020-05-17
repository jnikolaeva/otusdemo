CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL PRIMARY KEY,
    username VARCHAR(256) UNIQUE,
    first_name VARCHAR(256),
    last_name VARCHAR(256),
    email VARCHAR(256),
    phone VARCHAR(256)
);