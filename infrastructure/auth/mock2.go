package auth

import (
	"context"
	"kiravia.com/internship-go-api/domain"
)

type mock2 struct {
}

func (a mock2) VerifyToken(ctx context.Context, token string) (string, error) {
	return "mock2", nil
}

func NewAuthMock2Client() domain.AuthClient {
	return &mock2{}
}
