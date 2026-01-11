package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/services"
)

func ValidateAltcha(db queries.DBTX, hmacKey []byte) echo.MiddlewareFunc {
	service := services.NewAltchaService(db, hmacKey)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			payload := c.Request().FormValue("altcha")
			if payload == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "Missing ALTCHA payload")
			}

			_, err := service.VerifyPayload(c.Request().Context(), payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}

			return next(c)
		}
	}
}
