-- +goose Up
-- +goose StatementBegin
create table comments (
  id integer not null primary key,
  post_id integer not null references posts (id),
  body text not null,
  signature text not null,
  website text,
  created_at integer not null default (unixepoch()),
  updated_at integer not null default (unixepoch())
);

create index comments_post_id_index on comments (post_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table comments;
-- +goose StatementEnd
