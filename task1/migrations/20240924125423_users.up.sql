CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
)