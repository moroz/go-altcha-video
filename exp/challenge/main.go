package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/altcha-org/altcha-lib-go"
	"github.com/moroz/go-altcha-video/config"
)

func main() {
	expires := time.Now().Add(20 * time.Minute)
	challenge, err := altcha.CreateChallenge(altcha.ChallengeOptions{
		HMACKey:   string(config.AltchaChallengeSigner),
		MaxNumber: 500_000,
		Expires:   &expires,
	})
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	encoder.Encode(challenge)

	solution, err := altcha.SolveChallenge(challenge.Challenge, challenge.Salt, altcha.Algorithm(challenge.Algorithm), int(challenge.MaxNumber), 0, nil)

	if err != nil {
		log.Fatal(err)
	}

	encoder.Encode(solution)
}
