create table post
(
    id        integer,
    title     text(100)               not null,
    content   text(5000)              not null,
    author_id integer references user not null,
    primary key (id autoincrement)
);