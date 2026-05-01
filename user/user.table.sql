create table user
(
    id       integer,
    username text(100) not null,
    password text(100) not null,
    email    text(100) not null unique,
    primary key (id autoincrement)
);
