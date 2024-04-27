CREATE TABLE public.users (
	id serial4 NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	firebase_uid varchar(255) NOT NULL,
	firebase_provider_id varchar(255) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_uk UNIQUE (email),
	CONSTRAINT users_uk2 UNIQUE (firebase_uid)
);