-- +goose Up
-- +goose StatementBegin
create table posts (
  id integer not null primary key,
  title text not null,
  body text not null,
  slug text not null unique,
  published_at integer,
  created_at integer not null default (unixepoch()),
  updated_at integer not null default (unixepoch())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table posts;
-- +goose StatementEnd
