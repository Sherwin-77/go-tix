BEGIN;

CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    triggerable_id UUID NOT NULL,
    triggerable_type VARCHAR(255) NOT NULL,
    read_at TIMESTAMP(6) WITH TIME ZONE,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    metadata JSONB,

    CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

END;