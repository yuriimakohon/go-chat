CREATE TABLE IF NOT EXISTS users
(
    login         varchar(255) primary key,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS rooms
(
    id varchar(255) primary key
);

CREATE TABLE IF NOT EXISTS members
(
    user_login varchar(255) references users (login) on delete cascade not null,
    room_id    varchar(255) references rooms (id) on delete cascade    not null
);

CREATE TABLE IF NOT EXISTS messages
(
    id         serial                                                  not null unique,
    user_login varchar(255) references users (login) on delete cascade not null,
    room_id    varchar(255) references rooms (id) on delete cascade    not null,
    content    text                                                    not null,
    time       timestamp with time zone                                not null
);
