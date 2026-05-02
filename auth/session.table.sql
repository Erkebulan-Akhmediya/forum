create table session
(
    id         text(100) primary key,
    user_id    integer  not null unique references user,
    expires_at datetime not null
);
