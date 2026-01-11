package services

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/altcha-org/altcha-lib-go"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/internal/dbtypes"
)

type ChallengeService struct {
	queries *queries.Queries
	hmacKey string
}

func NewChallengeService(db queries.DBTX, hmacKey []byte) *ChallengeService {
	return &ChallengeService{
		queries: queries.New(db),
		hmacKey: string(hmacKey),
	}
}

var ErrChallengeReplay = errors.New("challenge solution has already been submitted")
var ErrInvalidSolution = errors.New("invalid solution")
var ErrNoExpirationTime = errors.New("expiration time not found in payload")
var ErrNoSolution = errors.New("no solution provided")

func (s *ChallengeService) ValidateChallenge(ctx context.Context, payload string) (bool, error) {
	if payload == "" {
		return false, ErrNoSolution
	}

	hash := sha256.Sum256([]byte(payload))
	if used, _ := s.queries.CheckUsedAltchaChallenge(ctx, hash[:]); used {
		return false, ErrChallengeReplay
	}

	ok, err := altcha.VerifySolutionSafe(payload, s.hmacKey, true)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, ErrInvalidSolution
	}

	expiration, err := extractExpirationTime(payload)
	if err != nil {
		return false, err
	}
	_, err = s.queries.InsertUsedAltchaChallenge(ctx, queries.InsertUsedAltchaChallengeParams{
		ChallengeHash: hash[:],
		ExpiresAt:     dbtypes.UnixTimestamp{Time: expiration},
	})

	return err == nil, err
}

func extractExpirationTime(payload string) (*time.Time, error) {
	var parsedPayload altcha.Payload
	bytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &parsedPayload)
	if err != nil {
		return nil, err
	}

	params := altcha.ExtractParams(parsedPayload)
	expires := params.Get("expires")
	if expires == "" {
		return nil, ErrNoExpirationTime
	}

	unixTimestamp, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return nil, ErrNoExpirationTime
	}

	ts := time.Unix(unixTimestamp, 0)
	return &ts, nil
}
