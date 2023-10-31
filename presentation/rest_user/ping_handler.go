package rest_user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PingHandler struct {
}

func (h PingHandler) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "PONG")
}
