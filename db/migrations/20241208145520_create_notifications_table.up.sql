BEGIN;

CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    triggerable VARCHAR(50),
    read_at DATE,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    metadata JSON,

    CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

END;