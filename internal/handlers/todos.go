package handlers

import (
	"github.com/gin-gonic/gin"
	"go-web-base/internal/services"
	"go-web-base/internal/types"
	"log"
	"net/http"
)

type TodoRequestHandler interface {
	PublicIndexPageHandler(ctx *gin.Context)
	PublicViewPageHandler(ctx *gin.Context)

	AppIndexPageHandler(ctx *gin.Context)
	AppViewPageHandler(ctx *gin.Context)
	AppNewPageHandler(ctx *gin.Context)
	AppEditPageHandler(ctx *gin.Context)

	AppNewActionHandler(ctx *gin.Context)
	AppEditActionHandler(ctx *gin.Context)
	AppDeleteActionHandler(ctx *gin.Context)
}

type TodoHandler struct {
	todoService *services.TodoService
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) FetcherMiddleware(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) != 0 {
		record, err := h.todoService.GetByID(ctx, id)
		if err != nil {
			log.Println(err)
		}

		ctx.Set("todo", record)
	}

	ctx.Next()
}

func (h *TodoHandler) PublicIndexPageHandler(ctx *gin.Context) {
	records, err := h.todoService.GetAll(ctx)
	if err != nil {
		log.Println("Handle Error")
	}

	ctx.HTML(http.StatusOK, "public/index", gin.H{
		"Records": records,
		"Meta": types.Meta{
			Title: "Todos",
		},
	})
}

func (h *TodoHandler) PublicViewPageHandler(ctx *gin.Context) {
	record := h.todoService.GetResourceFromContext(ctx)

	ctx.HTML(http.StatusOK, "public/todos/view", gin.H{
		"Record": record,
		"Meta": types.Meta{
			Title: "View Todo",
		},
	})
}

func (h *TodoHandler) AppIndexPageHandler(ctx *gin.Context) {
	records, err := h.todoService.GetAll(ctx)
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "app/todos/index", gin.H{
		"Records": records,
		"Meta": types.Meta{
			Title: "Todos",
		},
	})
}

func (h *TodoHandler) AppViewPageHandler(ctx *gin.Context) {
	record := h.todoService.GetResourceFromContext(ctx)

	ctx.HTML(http.StatusOK, "app/todos/view", gin.H{
		"Record": record,
		"Meta": types.Meta{
			Title: "View Todo",
		},
	})
}

func (h *TodoHandler) AppNewPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "app/todos/new", gin.H{
		"Meta": types.Meta{
			Title: "New Todo",
		},
	})
}

func (h *TodoHandler) AppEditPageHandler(ctx *gin.Context) {
	record := h.todoService.GetResourceFromContext(ctx)

	ctx.HTML(http.StatusOK, "app/todos/edit", gin.H{
		"Record": record,
		"Meta": types.Meta{
			Title: "Edit Todo",
		},
	})
}

/*
	Action Handlers
*/

func (h *TodoHandler) AppNewActionHandler(ctx *gin.Context) {
	var input types.NewTodoFormInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Println(err)
	}

	record, err := h.todoService.Create(ctx, input)
	if err != nil {
		ctx.HTML(http.StatusOK, "todos/fragments/forms/new", gin.H{
			"globalError": "Internal Server Error",
		})
		return
	}

	ctx.HTML(http.StatusOK, "fragments/todos/card", gin.H{
		"Body": record.Body,
	})
}

func (h *TodoHandler) AppEditActionHandler(ctx *gin.Context) {
	record := h.todoService.GetResourceFromContext(ctx)

	var input types.EditTodoFormInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Println(err)
	}

	_, err = h.todoService.Update(ctx, record, input)
	if err != nil {
		ctx.HTML(http.StatusOK, "todos/fragments/forms/edit", gin.H{
			"globalError": "Internal Server Error",
		})
		return
	}

	ctx.Header("HX-Redirect", "/app/todos")
}

func (h *TodoHandler) AppDeleteActionHandler(ctx *gin.Context) {
	record := h.todoService.GetResourceFromContext(ctx)

	err := h.todoService.Delete(ctx, record)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
}
