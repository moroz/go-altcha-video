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

	r.Static("/assets", "assets/static")

	blog := BlogController(db)
	r.GET("/", blog.Index)
	r.GET("/blog/:slug", blog.Show)

	altchaMiddleware := NewAltchaMiddleware(db, config.AltchaChallengeSigner)

	comments := CommentController(db)
	r.POST("/blog/:slug/comments", altchaMiddleware.Use(comments.Create))

	altcha := AltchaController()
	r.GET("/api/challenge", altcha.Challenge)

	return r
}
