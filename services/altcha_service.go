package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/altcha-org/altcha-lib-go"
)

type altchaService struct {
	hmacKey []byte
}

type AltchaService interface {
	VerifyPayload(payload string) (bool, error)
}

func NewAltchaService(hmacKey []byte) AltchaService {
	return &altchaService{
		hmacKey: hmacKey,
	}
}

var ErrInvalidSolution = errors.New("invalid solution to the challenge")

func (s *altchaService) VerifyPayload(payload string) (bool, error) {
	bytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return false, err
	}

	var parsedPayload altcha.Payload
	if err := json.Unmarshal(bytes, &parsedPayload); err != nil {
		return false, err
	}

	return altcha.VerifySolutionSafe(parsedPayload, string(s.hmacKey), true)
}
