BEGIN;

CREATE TABLE otps (
    id uuid SERIAL PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) NOT NULL,
    user_id INTEGER,
    action VARCHAR(50),
    code VARCHAR(10),
    status VARCHAR(20),
    sent_at DATETIME,
    verified_at DATETIME,
    expired_at DATETIME,
    verify_attempts INTEGER DEFAULT 0,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,
);

END;