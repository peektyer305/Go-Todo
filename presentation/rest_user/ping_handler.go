package rest_user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kiravia.com/internship-go-api/di"
	"net/http"
)

type PingHandler struct {
}

func (h PingHandler) Ping(ctx echo.Context) error {

	// Note: di.HogeHoge() で依存関係が解決されたUseCaseを取得できる
	result, err := di.SamplePingPong().Exec(ctx.Request().Context())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return ctx.String(http.StatusOK, result)
}
