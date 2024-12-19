BEGIN;

CREATE TABLE otps (
    id UUID PRIMARY KEY ,
    email VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    action VARCHAR(50),
    code VARCHAR(10),
    status VARCHAR(20),
    sent_at DATE,
    verified_at DATE,
    expired_at DATE,
    verify_attempts INTEGER,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT otp_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

END;