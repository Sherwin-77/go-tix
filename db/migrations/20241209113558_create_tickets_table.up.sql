BEGIN;

CREATE TABLE tickets (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL,
    category VARCHAR(100),
    price DECIMAL(12, 2),
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT tickets_event_id_fkey FOREIGN KEY (event_id) REFERENCES events (id)
);

END;