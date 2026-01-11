package handlers

import (
	"net/http"

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
		if _, err := me.ChallengeService.ValidateChallenge(c.Request().Context(), payload); err != nil {
			c.Logger().Errorf("AltchaMiddleware: %s", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Altcha payload")
		}

		return next(c)
	}
}
