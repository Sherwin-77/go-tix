BEGIN; 

CREATE TABLE event_approvals (
    id uuid PRIMARY KEY,
    user_id INTEGER NOT NULL ,
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
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT event_approvals_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);
END;