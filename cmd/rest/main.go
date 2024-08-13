package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peektyer305/Go-Todo/config"
	"github.com/peektyer305/Go-Todo/infrastructure"
	"github.com/peektyer305/Go-Todo/presentation/rest_todo"
)

func main() {
	// ToDo: DockerComposeを利用してdatabaseを作成することができたら、以下のコメントアウトを外す
	db := infrastructure.NewGormPostgres()
	defer func() {
		d, _ := db.DB()
		d.Close()
	}()

	engine := echo.New()
	engine.Debug = true

	engine.Pre(middleware.RemoveTrailingSlash())
	engine.Use(middleware.Recover())

	engine.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	TodoHandler := rest_todo.NewTodoHandler()
	todoGroup := engine.Group("/todos")
	rest_todo.RouteInit(todoGroup, TodoHandler)

	go func() {
		if err := engine.Start(fmt.Sprintf(":%s", config.Conf.GetPort())); err != nil && !errors.Is(err, http.ErrServerClosed) {
			engine.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := engine.Shutdown(ctx); err != nil {
		engine.Logger.Fatal(err)
	}
	println("stop server method")
}
