package types

import "github.com/moroz/go-altcha-video/db/queries"

type PostDetailsDto struct {
	*queries.Post
	Comments []*queries.Comment
}

type PostListDto struct {
	*queries.Post
	CommentCount int
}
