package handlers

import (
	"net/http"
	"time"

	"github.com/altcha-org/altcha-lib-go"
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/config"
)

type altchaController struct{}

func AltchaController() *altchaController {
	return &altchaController{}
}

func (*altchaController) Challenge(c echo.Context) error {
	expires := time.Now().Add(20 * time.Minute)

	challenge, err := altcha.CreateChallenge(altcha.ChallengeOptions{
		HMACKey:   string(config.AltchaChallengeSigner),
		MaxNumber: config.AltchaDifficulty,
		Expires:   &expires,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, challenge)
}
