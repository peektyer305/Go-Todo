package rest_todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/peektyer305/Go-Todo/application/todo"
	di "github.com/peektyer305/Go-Todo/di"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
	myError "github.com/peektyer305/Go-Todo/errors"
	request "github.com/peektyer305/Go-Todo/presentation/rest_todo/request"
	response "github.com/peektyer305/Go-Todo/presentation/rest_todo/response"
)

type TodoHandler struct {
	TodoUseCase *todo.TodoUseCase
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		TodoUseCase: di.InitializeTodoUseCase(),
	}
}
func (h *TodoHandler) FindAllByCriterias(ctx echo.Context) error{
	queryParams := ctx.QueryParams()
	var params request.FindParams
	if id, exists := queryParams["id"]; exists {
		valueobject.NewTodoId(id[0])
	}
	if title, exists := queryParams["title"]; exists {
		params.Title = &title[0]
	}
	if body, exists := queryParams["body"]; exists {
		params.Body = &body[0]
	}
	if isCompleted, exists := queryParams["isCompleted"]; exists {
		isCompletedBool := isCompleted[0] == "true"
		params.IsCompleted = &isCompletedBool
	}

	todos, err := h.TodoUseCase.FindAllByCriterias(ctx.Request().Context(), params)
	if err, ok := err.(*myError.NotFoundError); ok {
		return ctx.JSON(http.StatusNotFound, err)
	} 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	var responseTodos []response.TodoResponse
	for _, todo := range todos {
		response := response.CreateTodoResponse(
			todo.Id,
			todo.Title,
			todo.Body,
			todo.DueDate,
			todo.CompletedAt,
			todo.CreatedAt,
			todo.UpdatedAt,
		)
		responseTodos = append(responseTodos, response)
	}
	return ctx.JSON(http.StatusOK, responseTodos)
}

func (h *TodoHandler) FindById(ctx echo.Context) error{
	idParam := ctx.Param("id")
	id, err := valueobject.NewTodoId(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo, err := h.TodoUseCase.FindById(ctx.Request().Context(), id)
	if err, ok := err.(*myError.NotFoundError); ok {
		return ctx.JSON(http.StatusNotFound, err)
	} 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := response.CreateTodoResponse(
		todo.Id,
		todo.Title,
		todo.Body,
		todo.DueDate,
		todo.CompletedAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	return ctx.JSON(http.StatusOK, response)
	
}

func (h *TodoHandler) Create(ctx echo.Context) error{
	var params request.CreateParams
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := h.TodoUseCase.Create(ctx.Request().Context(), params)
	
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	
	response := response.CreateTodoResponse(
		todo.Id,
		todo.Title,
		todo.Body,
		todo.DueDate,
		todo.CompletedAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	return ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) UpdateById(ctx echo.Context) error{
	var params request.UpdateParams
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	id := ctx.Param("id")
	todoId,err := valueobject.NewTodoId(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todo,err := h.TodoUseCase.UpdateById(ctx.Request().Context(), todoId, params)
	if err, ok := err.(*myError.NotFoundError); ok {
		return ctx.JSON(http.StatusNotFound, err)
	} 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := response.CreateTodoResponse(
		todo.Id,
		todo.Title,
		todo.Body,
		todo.DueDate,
		todo.CompletedAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	return ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) DeleteById(ctx echo.Context) error{
	id := ctx.Param("id")
	todoId,err := valueobject.NewTodoId(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	err = h.TodoUseCase.DeleteById(ctx.Request().Context(), todoId)
	if err, ok := err.(*myError.NotFoundError); ok {
		return ctx.JSON(http.StatusNotFound, err)
	} 
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
	todo,err := h.TodoUseCase.CopyById(ctx.Request().Context(), todoId)
	if err, ok := err.(*myError.NotFoundError); ok {
		return ctx.JSON(http.StatusNotFound, err)
	} 
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := response.CreateTodoResponse(
		todo.Id,
		todo.Title,
		todo.Body,
		todo.DueDate,
		todo.CompletedAt,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	return ctx.JSON(http.StatusOK, response)
}

func RouteInit(routeGroup *echo.Group, handler *TodoHandler) {
	todo := routeGroup.Group("/todos")
	todo.GET("/search", handler.FindAllByCriterias)
	todo.GET("/:id", handler.FindById)
	todo.POST("", handler.Create)
	todo.POST("/:id/copy", handler.CopyById)
	todo.DELETE("/:id/delete",handler.DeleteById)
	todo.PATCH("/:id/update", handler.UpdateById)
}