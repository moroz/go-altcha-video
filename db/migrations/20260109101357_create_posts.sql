-- +goose Up
-- +goose StatementBegin
create table posts (
  id uuid not null primary key,
  author text not null default 'Karol Moroz',
  title text not null,
  body text not null,
  slug text not null unique,
  published_at unix_timestamp,
  created_at unix_timestamp not null default (unixepoch()),
  updated_at unix_timestamp not null default (unixepoch())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table posts;
-- +goose StatementEnd
