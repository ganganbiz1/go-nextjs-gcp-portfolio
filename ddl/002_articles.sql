CREATE TABLE public.articles (
	id serial4 NOT NULL,
	title varchar(255) NOT NULL,
	content text NOT NULL,
	image_id int4 NOT NULL,
	user_id int4 NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	CONSTRAINT articles_pk PRIMARY KEY (id)
);