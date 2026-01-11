package config

import (
	"encoding/base64"
	"log"
	"time"
)

func MustGetenvBase64(name string) []byte {
	val := MustGetenv(name)
	bytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		log.Fatalf("Failed to decode environment variable %s from Base64: %s", name, err)
	}
	return bytes
}

const AltchaDifficulty = 500_000
const AltchaChallengeExpirationTime = 20 * time.Minute

var AltchaHMACKey = MustGetenvBase64("ALTCHA_HMAC_KEY")
