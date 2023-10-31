package di

import (
	"github.com/google/wire"
	"kiravia.com/internship-go-api/infrastructure"
	"kiravia.com/internship-go-api/infrastructure/auth"
)

var providerSet = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// client
	auth.NewAuthMockClient,

	// Repository

	// queryService

	// domainService

	// useCase
)
