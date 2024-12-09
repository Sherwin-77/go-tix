CREATE TABLE event_approvals (
    id uuid SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    title VARCHAR(255) NOT NULL,
    description VARCHAR(2047),
    organizer VARCHAR(255) NOT NULL,
    location VARCHAR(2047),
    longitude DECIMAL(12, 2),
    latitude DECIMAL(12, 2),
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);
END;