-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table "public"."samples"
(
    "id"         uuid        not null,
    "name"       text        not null unique,
    "created_at" timestamptz not null,
    "updated_at" timestamptz not null,
    primary key ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table "public"."samples";
-- +goose StatementEnd