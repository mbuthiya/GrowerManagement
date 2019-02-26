create table Users(
    id serial primary key,
    email varchar(255) not null,
    password text
);