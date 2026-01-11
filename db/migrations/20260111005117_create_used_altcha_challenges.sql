-- +goose Up
-- +goose StatementBegin
create table used_altcha_challenges (
  id integer not null primary key,
  challenge_hash blob not null unique,
  expires_at unix_timestamp not null,
  created_at unix_timestamp not null default (unixepoch())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table used_altcha_challenges;
-- +goose StatementEnd
