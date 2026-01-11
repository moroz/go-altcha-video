package config

import (
	"crypto/hkdf"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

func GetenvWithDefaultBase64(name, defaultValue string) []byte {
	val := GetenvWithDefault(name, defaultValue)
	bytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		log.Fatalf("Failed to decode environment variable as Base64: %s", err)
	}
	return bytes
}

func MustDeriveSecret(info string, length int) []byte {
	derived, err := hkdf.Key(sha256.New, SecretKeyBase, []byte(HkdfSalt), info, length)
	if err != nil {
		log.Fatalf("Failed to derive key: %s", err)
	}
	return derived
}

const HkdfSalt = "e1c65e55e6b0ba0e83619cf114fce86c0e48b52d5d9b6a5c5f7c119d55d8eb2e653d3269ed0ee7ceb96648613840bf9b"

var SecretKeyBase = GetenvWithDefaultBase64("SECRET_KEY_BASE", "75agdtxHgBf4sDDia0UKNISrTmy1jM7BqJJu1zbq++N+Ag4NqQ/+eahMmP7oSNGPhGzVg4A+T+eOwBnOC9iqKg==")
var AltchaChallengeSigner = MustDeriveSecret("ALTCHA challenge signer", 32)

const AltchaDifficulty = 100_000
