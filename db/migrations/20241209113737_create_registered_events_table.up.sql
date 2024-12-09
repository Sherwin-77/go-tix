BEGIN;

CREATE TABLE registered_events (
    id uuid SERIAL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255),
    description TEXT,
    price DECIMAL(12, 2),
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;