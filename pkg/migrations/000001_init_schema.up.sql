CREATE TABLE IF NOT EXISTS user_role
(
    id   integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    role text NOT NULL
);

INSERT INTO user_role (role)
values ('ADMIN'),
       ('USER');

CREATE TABLE IF NOT EXISTS users
(
    id       integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    login    text    NOT NULL,
    password text    NOT NULL,
    roleID   integer NOT NULL REFERENCES user_role (id)
);

INSERT INTO users (login, password, roleID)
values ('ADMIN', 'ADMIN', 1);

CREATE TABLE IF NOT EXISTS purchase
(
    id       integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    userID   integer NOT NULL REFERENCES users (id),
    date     date    NOT NULL,
    fileName text    NOT NULL
);

CREATE TABLE IF NOT EXISTS comment
(
    id         integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    userID     integer NOT NULL REFERENCES users (id),
    purchaseID integer NOT NULL REFERENCES purchase (id),
    date       date    NOT NULL,
    text       text    NOT NULL
);
