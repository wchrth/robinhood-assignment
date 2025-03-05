CREATE TABLE comments (
    id bigserial NOT NULL,
    description text NOT NULL,
    user_id bigint NOT NULL,
    appointment_id bigint NOT NULL,
    created_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(100) NOT NULL,
    updated_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100) NOT NULL,
    CONSTRAINT comments_pkey PRIMARY KEY (id),
    CONSTRAINT comments_fkey_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT comments_fkey_appointment FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE CASCADE
);
