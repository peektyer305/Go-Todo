package sample

import (
	"context"
	"fmt"
	"kiravia.com/internship-go-api/domain"
	"time"
)

type PingPong struct {
	auth domain.AuthClient
}

// Exec
// Note: Goではctxを引数にとって伝番していく
func (s *PingPong) Exec(ctx context.Context) (string, error) {
	uid, err := s.auth.VerifyToken(ctx, "token")
	if err != nil {
		return "", fmt.Errorf("failed to verify token: %w", err)
	}
	println(uid)

	return fmt.Sprintf("pong %s", time.Now().String()), nil
}

// NewSamplePingPong
// Note: 本来 domain.AuthClient は不要だが、DIのサンプルのために入れている
func NewSamplePingPong(auth domain.AuthClient) *PingPong {
	return &PingPong{
		auth: auth,
	}
}
