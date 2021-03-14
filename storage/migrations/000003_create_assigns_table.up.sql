-- migrate create -ext sql -dir ./storage/migrations -seq create_assigns_table
CREATE TABLE IF NOT EXISTS assigns (
    id serial PRIMARY KEY,
    ticket_id integer NOT NULL,
    user_id integer NOT NULL,
    UNIQUE(ticket_id, user_id),
    constraint fk_ticket_id foreign key (ticket_id) REFERENCES tickets (id),
    constraint fk_user_id foreign key (user_id) REFERENCES users (id)
);