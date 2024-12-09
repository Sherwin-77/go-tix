BEGIN;

CREATE TABLE notifications (
    id uuid INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    triggerable VARCHAR(50),
    read_at DATE,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    metadata JSON
);

END;