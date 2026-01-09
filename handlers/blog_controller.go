package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/services"
	"github.com/moroz/go-altcha-video/tmpl/blog"
)

type blogController struct {
	PostService *services.PostService
}

func BlogController(db queries.DBTX) *blogController {
	return &blogController{
		PostService: services.NewPostService(db),
	}
}

func (me *blogController) Index(c echo.Context) error {
	posts, err := me.PostService.ListPosts(c.Request().Context())
	if err != nil {
		return err
	}

	return blog.Index(posts).Render(c.Response().Writer)
}

func (me *blogController) Show(c echo.Context) error {
	post, err := me.PostService.GetPostDetailsBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil {
		return err
	}

	return blog.Show(post).Render(c.Response().Writer)
}
