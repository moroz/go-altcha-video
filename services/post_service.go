package services

import (
	"context"

	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/types"
)

type PostService struct {
	queries *queries.Queries
}

func NewPostService(db queries.DBTX) *PostService {
	return &PostService{
		queries: queries.New(db),
	}
}

func (s *PostService) GetPostDetailsBySlug(ctx context.Context, slug string) (*types.PostDetailsDto, error) {
	post, err := s.queries.GetPostBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	comments, err := s.queries.GetCommentsByPostId(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	return &types.PostDetailsDto{
		Post:     post,
		Comments: comments,
	}, nil
}

func (s *PostService) ListPosts(ctx context.Context) ([]*types.PostListDto, error) {
	posts, err := s.queries.ListPosts(ctx)
	if err != nil {
		return nil, err
	}

	commentCounts, err := s.GetCommentCountsForPosts(ctx, posts)
	if err != nil {
		return nil, err
	}

	result := make([]*types.PostListDto, len(posts))
	for i, post := range posts {
		result[i] = &types.PostListDto{
			Post:         post,
			CommentCount: commentCounts[post.ID],
		}
	}

	return result, nil
}

func (s *PostService) GetCommentCountsForPosts(ctx context.Context, posts []*queries.Post) (map[int64]int, error) {
	result := make(map[int64]int)

	if len(posts) == 0 {
		return result, nil
	}

	ids := make([]int64, len(posts))
	for i, post := range posts {
		ids[i] = post.ID
	}

	counts, err := s.queries.GetCommentCountsForPosts(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, count := range counts {
		result[count.PostID] = int(count.Count)
	}

	return result, nil
}
