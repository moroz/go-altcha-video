-- name: GetCommentsByPostId :many
select * from comments where post_id = ? order by id desc;

-- name: GetCommentCountsForPosts :many
select c.post_id, count(c.id) from comments c group by 1;

-- name: InsertComment :one
insert into comments (id, post_id, signature, body, website)
values (?, ?, ?, ?, ?)
returning *;
