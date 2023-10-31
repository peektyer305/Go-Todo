//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"kiravia.com/internship-go-api/application/sample"
	"kiravia.com/internship-go-api/infrastructure"
	"kiravia.com/internship-go-api/infrastructure/auth"
)

var providerSet = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// client
	auth.NewAuthMockClient,
	// Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	//auth.NewAuthMock2Client,

	// Repository

	// queryService

	// domainService

	// useCase
	sample.NewSamplePingPong,
)

func SamplePingPong() *sample.PingPong {
	wire.Build(
		providerSet,
	)
	return nil
}
