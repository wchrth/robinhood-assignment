CREATE TABLE users (
	id bigserial NOT NULL,
	email varchar(100) NOT NULL,
	"password" varchar(100) NOT NULL,
	display_name varchar(100) NOT NULL,
	created_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_email_key UNIQUE (email)
);
