BEGIN;

CREATE TABLE role_users (
    role_id UUID NOT NULL,
    user_id UUID NOT NULL,

    CONSTRAINT role_users_pkey PRIMARY KEY (role_id, user_id),
    CONSTRAINT role_users_role_id_fkey FOREIGN KEY (role_id) REFERENCES roles(id),
    CONSTRAINT role_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

COMMIT;