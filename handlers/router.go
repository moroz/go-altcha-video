package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
)

func Router(db queries.DBTX) *echo.Echo {
	r := echo.New()

	blog := BlogController(db)
	r.GET("/", blog.Index)

	return r
}
