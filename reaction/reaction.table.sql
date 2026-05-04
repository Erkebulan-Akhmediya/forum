create table reaction
(
    id         integer,
    user_id    integer references user not null,
    post_id    integer references post,
    comment_id integer references comment,
    reaction   text(10) check ( reaction in ('LIKE', 'DISLIKE') ) not null,
    primary key (id autoincrement)
);