CREATE TABLE IF NOT EXISTS users
(
    id       integer                           NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    login    text  NOT NULL,
    password text NOT NULL
);
