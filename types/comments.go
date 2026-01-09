package types

type CreateCommentParams struct {
	Signature string  `schema:"signature"`
	Body      string  `schema:"body"`
	Website   *string `schema:"website"`
}
