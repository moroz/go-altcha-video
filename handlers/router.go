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
	r.File("/assets/app.mjs", "assets/app.mjs")

	blog := BlogController(db)
	r.GET("/", blog.Index)
	r.GET("/blog/:slug", blog.Show)

	altcha := AltchaController(config.AltchaHMACKey)
	r.GET("/api/v1/challenge", altcha.Challenge)

	comments := CommentController(db)
	g := r.Group("")
	g.Use(ValidateAltcha(config.AltchaHMACKey))
	g.POST("/blog/:slug/comments", comments.Create)

	return r
}
