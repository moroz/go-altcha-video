-- name: ListPosts :many
select * from posts
where published_at is not null
order by published_at desc;

-- name: GetPostBySlug :one
select * from posts
where slug = ?;
