package handlers

import (
	"net/http"
	"time"

	"github.com/altcha-org/altcha-lib-go"
	"github.com/labstack/echo/v4"
	"github.com/moroz/go-altcha-video/config"
)

type altchaController struct {
	hmacKey []byte
}

func AltchaController(hmacKey []byte) *altchaController {
	return &altchaController{
		hmacKey: hmacKey,
	}
}

func (me *altchaController) Challenge(c echo.Context) error {
	expiresAt := time.Now().Add(config.AltchaChallengeExpirationTime)

	challenge, err := altcha.CreateChallenge(altcha.ChallengeOptions{
		HMACKey:   string(me.hmacKey),
		MaxNumber: config.AltchaDifficulty,
		Expires:   &expiresAt,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, challenge)
}
