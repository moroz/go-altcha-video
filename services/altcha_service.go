package services

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/altcha-org/altcha-lib-go"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/internal/dbtypes"
)

type altchaService struct {
	hmacKey []byte
	queries *queries.Queries
}

type AltchaService interface {
	VerifyPayload(ctx context.Context, payload string) (bool, error)
}

func NewAltchaService(db queries.DBTX, hmacKey []byte) AltchaService {
	return &altchaService{
		hmacKey: hmacKey,
		queries: queries.New(db),
	}
}

var ErrInvalidSolution = errors.New("invalid solution to the challenge")

func (s *altchaService) VerifyPayload(ctx context.Context, payload string) (bool, error) {
	bytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return false, err
	}

	var parsedPayload altcha.Payload
	if err := json.Unmarshal(bytes, &parsedPayload); err != nil {
		return false, err
	}

	hash := sha256.Sum256([]byte(payload))
	reused, err := s.queries.ValidateChallengeReuse(ctx, hash[:])
	if reused || err != nil {
		return false, ErrInvalidSolution
	}

	ok, err := altcha.VerifySolutionSafe(parsedPayload, string(s.hmacKey), true)
	if err != nil {
		return false, ErrInvalidSolution
	}

	expiresAt, err := extractExpiresAt(parsedPayload)
	if err != nil {
		return false, err
	}

	_, err = s.queries.InsertUsedChallenge(ctx, queries.InsertUsedChallengeParams{
		ChallengeHash: hash[:],
		ExpiresAt: dbtypes.UnixTimestamp{
			Time: expiresAt,
		},
	})
	if err != nil {
		return false, err
	}

	return ok, nil
}

func extractExpiresAt(payload altcha.Payload) (*time.Time, error) {
	params := altcha.ExtractParams(payload)
	expires := params.Get("expires")
	if expires == "" {
		return nil, errors.New("expiration time not set in the challenge")
	}

	unix, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("extractExpiresAt: %w", err)
	}

	ts := time.Unix(unix, 0)
	return &ts, nil
}
