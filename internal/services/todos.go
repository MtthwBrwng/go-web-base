package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/r3labs/diff"
	"go-web-base/internal/stores"
	"go-web-base/internal/types"
	"time"
)

type TodoService struct {
	todoStore *stores.TodoStore
}

func NewTodoService(todoStore *stores.TodoStore) *TodoService {
	return &TodoService{todoStore: todoStore}
}

func (s *TodoService) GetResourceFromContext(ctx *gin.Context) *types.Todo {
	resource, exists := ctx.Get("todo")
	if exists {
		return interface{}(resource).(*types.Todo)
	}

	return nil
}

func (s *TodoService) Create(ctx *gin.Context, input types.NewTodoFormInput) (*types.Todo, error) {
	newTodo := types.Todo{
		ID:      uuid.NewString(),
		Created: time.Now().UTC(),
		Updated: time.Now().UTC(),
		Body:    input.Body,
	}

	return s.todoStore.Create(ctx, newTodo)
}

func (s *TodoService) Update(ctx *gin.Context, record *types.Todo, input types.EditTodoFormInput) (*types.Todo, error) {
	updatedTodo := types.Todo{
		ID:      record.ID,
		Created: record.Created,
		Updated: record.Updated,
		Body:    input.Body,
	}

	if diff.Changed(record, updatedTodo) {
		updatedTodo.Updated = time.Now().UTC()
		return s.todoStore.Update(ctx, record.ID, updatedTodo)
	}

	return nil, nil
}

func (s *TodoService) Delete(ctx *gin.Context, record *types.Todo) error {
	return s.todoStore.Delete(ctx, record.ID)
}

func (s *TodoService) GetByID(ctx *gin.Context, uid string) (*types.Todo, error) {
	return s.todoStore.GetByID(ctx, uid)
}

func (s *TodoService) GetAll(ctx *gin.Context) ([]*types.Todo, error) {
	return s.todoStore.GetAll(ctx)
}
