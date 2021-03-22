-- migrate create -ext sql -dir ./storage/migrations -seq create_tags_tickets_table
CREATE TABLE IF NOT EXISTS tags_tickets (
    id serial PRIMARY KEY,
    ticket_id integer NOT NULL,
    tag_id integer NOT NULL,
    UNIQUE(ticket_id, tag_id),
    constraint fk_ticket_id foreign key (ticket_id) REFERENCES tickets (id),
    constraint fk_tag_id foreign key (tag_id) REFERENCES tags (id)
);