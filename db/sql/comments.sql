-- name: GetCommentsByPostId :many
select * from comments where post_id = ? order by created_at desc;

-- name: GetCommentCountsForPosts :many
select c.post_id, count(c.id) from comments c
where c.post_id = any(?)
group by 1;
