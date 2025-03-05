CREATE TABLE appointments (
    id bigserial NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    status varchar(50) NOT NULL,
    is_archived boolean NOT NULL,
    user_id bigint NOT NULL,
    created_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(100) NOT NULL,
    updated_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100) NOT NULL,
    CONSTRAINT appointments_pkey PRIMARY KEY (id),
    CONSTRAINT appointments_fkey_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
