create table if not exists accounts (
    id serial unique,
    name varchar(255) not null,
    password_hash varchar(255) not null,
    phone varchar(255) not null unique,
    description text
);