-- migrate create -ext sql -dir ./storage/migrations -seq create_tags_table
CREATE TABLE IF NOT EXISTS tags (
    id serial PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL
);