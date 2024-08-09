package rest_todo

import (
	"net/http"

	"github.com/labstack/echo"
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
	
	return ctx.String(http.StatusOK, "Create Todo")
}

func (h *TodoHandler) UpdateById(ctx echo.Context) error{
	
	return ctx.String(http.StatusOK, "Update Todo")
}

func (h *TodoHandler) DeleteById(ctx echo.Context) error{
	
	return ctx.String(http.StatusOK, "Delete Todo")
}

func (h *TodoHandler) CopyById(ctx echo.Context) error{
	
	return ctx.String(http.StatusOK, "Copy Todo")
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