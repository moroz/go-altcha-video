package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/types"
)

type CommentService struct {
	queries *queries.Queries
}

func NewCommentService(db queries.DBTX) *CommentService {
	return &CommentService{
		queries: queries.New(db),
	}
}

func (s *CommentService) CreateComment(ctx context.Context, post *queries.Post, params *types.CreateCommentParams) (*queries.Comment, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return s.queries.InsertComment(ctx, queries.InsertCommentParams{
		ID:        id,
		PostID:    post.ID,
		Signature: params.Signature,
		Body:      params.Body,
		Website:   params.Website,
	})
}
