CREATE TABLE public.accounts (
    id varchar(36) NOT NULL PRIMARY KEY,
    account_id varchar(36) NOT NULL
);
CREATE UNIQUE INDEX accounts_idx_account_id on public.accounts (account_id);

CREATE TABLE public.posts (
    id varchar(36) NOT NULL PRIMARY KEY,
    author_id varchar(36) NOT NULL,
    body text NOT NULL,
    CONSTRAINT posts_fk_author_id FOREIGN KEY (author_id) REFERENCES public.accounts (id)
);
CREATE INDEX posts_idx_author_id on public.posts (author_id);

CREATE TABLE public.comments (
    id varchar(36) NOT NULL PRIMARY KEY,
    post_id varchar(36) NOT NULL,
    author_id varchar(36) NOT NULL,
    body text NOT NULL,
    CONSTRAINT comments_fk_author_id FOREIGN KEY (author_id) REFERENCES public.accounts (id),
    CONSTRAINT comments_fk_posts_id FOREIGN KEY (post_id) REFERENCES public.posts (id)
);
CREATE INDEX comments_idx_post_id on public.comments (post_id);
CREATE INDEX comments_idx_author_id on public.comments (author_id);

CREATE TABLE public.relationships (
    id varchar(36) NOT NULL PRIMARY KEY,
    follower_id varchar(36) NOT NULL,
    followee_id varchar(36) NOT NULL,
    CONSTRAINT relationships_fk_follower_id FOREIGN KEY (follower_id) REFERENCES public.accounts (id),
    CONSTRAINT relationships_fk_followee_id FOREIGN KEY (followee_id) REFERENCES public.accounts (id)
);
CREATE UNIQUE INDEX relationships_idx_follower_followee_ids on public.relationships (follower_id, followee_id);
