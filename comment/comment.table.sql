create table comment
(
    id         integer,
    content    text(500)               not null,
    author_id  integer references user not null,
    post_id    integer references post,
    comment_id integer references comment,
    primary key (id autoincrement)
);