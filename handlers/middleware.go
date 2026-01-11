package handlers

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/services"
)

type altchaMiddleware struct {
	ChallengeService *services.ChallengeService
}

func NewAltchaMiddleware(db queries.DBTX, hmacKey []byte) *altchaMiddleware {
	return &altchaMiddleware{
		ChallengeService: services.NewChallengeService(db, hmacKey),
	}
}

func (me *altchaMiddleware) Use(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := c.FormValue("altcha")
		log.Print(payload)
		if _, err := me.ChallengeService.ValidateChallenge(c.Request().Context(), payload); err != nil {
			return fmt.Errorf("AltchaMiddleware: %w", err)
		}

		return next(c)
	}
}
