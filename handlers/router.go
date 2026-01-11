package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moroz/go-altcha-video/config"
	"github.com/moroz/go-altcha-video/db/queries"
)

func Router(db queries.DBTX) *echo.Echo {
	r := echo.New()

	r.Use(middleware.RequestLogger())

	r.File("/assets/output.css", "assets/output.css")

	blog := BlogController(db)
	r.GET("/", blog.Index)
	r.GET("/blog/:slug", blog.Show)

	comments := CommentController(db)
	r.POST("/blog/:slug/comments", comments.Create)

	altcha := AltchaController(config.AltchaHMACKey)
	r.GET("/api/v1/challenge", altcha.Challenge)

	return r
}
