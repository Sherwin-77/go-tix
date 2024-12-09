BEGIN;

CREATE TABLE notifications (
    id uuid SERIAL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    triggerable VARCHAR(50),
    read_at DATETIME,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    metadata JSON
);

END;