BEGIN;

CREATE TABLE tickets (
    id uuid SERIAL PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    category VARCHAR(100),
    price DECIMAL(12, 2),
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;