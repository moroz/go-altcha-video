package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/services"
	"github.com/moroz/go-altcha-video/types"
)

var formDecoder = schema.NewDecoder()

type commentController struct {
	PostService    *services.PostService
	CommentService *services.CommentService
}

func CommentController(db queries.DBTX) *commentController {
	return &commentController{
		PostService:    services.NewPostService(db),
		CommentService: services.NewCommentService(db),
	}
}

func (me *commentController) Create(c echo.Context) error {
	form, err := c.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var params types.CreateCommentParams
	err = formDecoder.Decode(&params, form)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	post, err := me.PostService.GetPostDetailsBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	_, err = me.CommentService.CreateComment(c.Request().Context(), post.Post, &params)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/blog/%s", post.Slug))
}
