-- migrate create -ext sql -dir ./storage/migrations -seq create_users_table
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL
);