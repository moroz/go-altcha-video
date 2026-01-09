-- +goose Up
-- +goose StatementBegin
create table comments (
  id uuid not null primary key,
  post_id uuid not null references posts (id) on delete cascade,
  body text not null,
  signature text not null,
  website text,
  created_at unix_timestamp not null default (unixepoch()),
  updated_at unix_timestamp not null default (unixepoch())
);

create index comments_post_id_index on comments (post_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table comments;
-- +goose StatementEnd
