CREATE TABLE public.accounts (
    id character(36) NOT NULL
);
ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);

CREATE TABLE public.posts (
    id character(36) NOT NULL,
    author_id character(36)
);
ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);

CREATE TABLE public.comments (
    id character(36) NOT NULL,
    author_id character varying(36)
);
ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);

CREATE TABLE public.relationships (
    id character(36) NOT NULL
);
ALTER TABLE ONLY public.relationships
    ADD CONSTRAINT relationships_pkey PRIMARY KEY (id);
