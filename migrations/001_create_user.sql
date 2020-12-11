-- tHis is a sample migration.

CREATE SEQUENCE user_seq AS integer;
CREATE TABLE user (
    id            integer PRIMARY KEY NOT NULL DEFAULT nextval('user_seq'),
    email         text UNIQUE,
    password_hash text,
    first_name    text,
    last_name     text,
    is_active     boolean NOT NULL DEFAULT TRUE,
    updated_at    TIMESTAMP NOT NULL DEFAULT now(),
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
ALTER SEQUENCE user_seq OWNED BY user.id;

---- create above / drop below ----

DROP table user;
