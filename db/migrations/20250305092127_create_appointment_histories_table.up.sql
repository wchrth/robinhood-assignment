CREATE TABLE appointment_histories (
    id bigserial NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    status varchar(50) NOT NULL,
    appointment_id bigint NOT NULL,
    CONSTRAINT appointment_histories_pkey PRIMARY KEY (id),
    CONSTRAINT appointment_histories_fkey_appointment FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE CASCADE
);
