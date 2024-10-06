create table users
(
    id       BIGSERIAL PRIMARY KEY,
    email    varchar(255) not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null
);
