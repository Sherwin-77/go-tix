BEGIN;

CREATE TABLE event_approval_tickets (
    id uuid PRIMARY KEY,
    event_approval_id INTEGER NOT NULL REFERENCES event_approvals(id) ON DELETE CASCADE,
    category VARCHAR(100),
    price DECIMAL(12, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT event_approval_tickets_event_approval_id_fkey FOREIGN KEY (event_approval_id) REFERENCES event_approvals (id)
);

END;