package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/services"
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
	return c.HTML(200, "<h1>Hello world!</h1>")
}
