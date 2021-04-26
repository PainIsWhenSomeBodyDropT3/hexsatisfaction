-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id       integer                           NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    login    text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL
) TABLESPACE pg_default;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF NOT EXISTS users;
-- +goose StatementEnd
