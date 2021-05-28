CREATE TABLE IF NOT EXISTS user_role
(
    id   integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    role text NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id       integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    login    text    NOT NULL,
    password text    NOT NULL,
    roleID   integer NOT NULL REFERENCES user_role (id)
);


INSERT INTO user_role (role)
values ('ADMIN'),
       ('USER');


INSERT INTO users (login, password, roleID)
values ('ADMIN', 'ADMIN', 1);

CREATE TABLE IF NOT EXISTS author
(
    id          integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name        text    NOT NULL,
    age         int     NOT NULL,
    description text    NOT NULL,
    userID      integer NOT NULL REFERENCES users (id)
);


CREATE TABLE IF NOT EXISTS file
(
    id          integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name        text    NOT NULL,
    description text    NOT NULL,
    size        integer NOT NULL,
    path        text    NOT NULL,
    addDate     date    NOT NULL,
    updateDate  date    NOT NULL,
    actual      boolean NOT NULL,
    authorID    integer NOT NULL REFERENCES author (id)
);


CREATE TABLE IF NOT EXISTS purchase
(
    id     integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    userID integer NOT NULL REFERENCES users (id),
    date   date    NOT NULL,
    fileID integer NOT NULL REFERENCES file (id)
);

CREATE TABLE IF NOT EXISTS comment
(
    id         integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    userID     integer NOT NULL REFERENCES users (id),
    purchaseID integer NOT NULL REFERENCES purchase (id),
    date       date    NOT NULL,
    text       text    NOT NULL
);
