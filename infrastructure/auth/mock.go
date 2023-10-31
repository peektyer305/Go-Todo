package auth

import (
	"context"
	"kiravia.com/internship-go-api/domain"
)

type mock struct {
}

func (a mock) VerifyToken(ctx context.Context, token string) (string, error) {
	return token, nil
}

func NewAuthMockClient() domain.AuthClient {
	return &mock{}
}
