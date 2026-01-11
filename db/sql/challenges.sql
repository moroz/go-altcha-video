-- name: ValidateChallengeReuse :one
select cast(exists (select 1 from used_altcha_challenges where challenge_hash = ?) as boolean);

-- name: InsertUsedChallenge :one
insert into used_altcha_challenges (challenge_hash, expires_at) values (?, ?) returning *;

-- name: VacuumUsedChallenges :exec
delete from used_altcha_challenges where expires_at < unixepoch();
