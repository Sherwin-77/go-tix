BEGIN;

CREATE TABLE events (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(2047),
    organizer VARCHAR(255) NOT NULL,
    location VARCHAR(2047),
    longitude DECIMAL(8, 6),
    latitude DECIMAL(9, 6),
    start_at DATE NOT NULL,
    end_at DATE NOT NULL,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;