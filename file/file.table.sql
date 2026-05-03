create table file
(
    id         integer,
    name       text(100) not null,
    post_id    integer references post,
    comment_id integer references comment unique,
    primary key (id autoincrement)
);
