CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp not null DEFAULT NOW(),
    unique (email)
);