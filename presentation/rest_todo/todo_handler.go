package rest_todo

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/peektyer305/Go-Todo/domain/entity"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type TodoHandler struct{

}
func (h *TodoHandler) FindAllByCriterias(ctx echo.Context) error{
	
	return ctx.String(http.StatusOK, "Search Todos")
}

func (h TodoHandler) FindById(ctx echo.Context) error{
	idParam := ctx.Param("id")
	id, err := valueobject.NewTodoId(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := di.TodoUseCase.FindById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
	
}

func (h *TodoHandler) Create(ctx echo.Context) error{
	var params entity.CreateParams
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := di.TodoUseCase.Create(ctx.Request().Context(), params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateById(ctx echo.Context) error{
	var params entity.UpdateParams
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	id := ctx.Param("id")
	todoId, err := valueobject.NewTodoId(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := di.TodoUseCase.UpdateById(ctx.Request().Context(), todoId, params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteById(ctx echo.Context) error{
	id := ctx.Param("id")
	todoId, err := valueobject.NewTodoId(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err = di.TodoUseCase.DeleteById(ctx.Request().Context(), todoId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "Todo Deleted")
}

func (h *TodoHandler) CopyById(ctx echo.Context) error{
	id := ctx.Param("id")
	todoId, err := valueobject.NewTodoId(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := di.TodoUseCase.CopyById(ctx.Request().Context(), todoId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, todo)
}

func RouteInit(routeGroup *echo.Group) {
	todo := routeGroup.Group("/todos")
	todo.GET("/search", (&TodoHandler{}).FindAllByCriterias)
	todo.GET("/:id", (&TodoHandler{}).FindById)
	todo.POST("/", (&TodoHandler{}).Create)
	todo.POST("/:id/copy", (&TodoHandler{}).CopyById)
	todo.DELETE("/:id/delete",(&TodoHandler{}).DeleteById)
	todo.PATCH("/:id/update", (&TodoHandler{}).UpdateById)
}