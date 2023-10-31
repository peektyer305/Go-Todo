package rest_user

import "github.com/labstack/echo/v4"

func RouteInit(routeGroup *echo.Group) {

	user := routeGroup.Group("/user")

	ping := user.Group("/ping")
	{
		ping.GET("", PingHandler{}.Ping)
	}

}
