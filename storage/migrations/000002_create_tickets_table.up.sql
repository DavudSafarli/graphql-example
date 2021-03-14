-- migrate create -ext sql -dir ./storage/migrations -seq create_tickets_table
CREATE TABLE IF NOT EXISTS tickets (
    id serial PRIMARY KEY,
    title VARCHAR (50) UNIQUE NOT NULL
);