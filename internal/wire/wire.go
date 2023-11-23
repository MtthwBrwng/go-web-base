//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"go-web-base/internal/handlers"
	"go-web-base/internal/services"
	"go-web-base/internal/stores"
	"go-web-base/internal/types"
)

func InitializeTodoHandler(store map[string]*types.Todo) *handlers.TodoHandler {
	wire.Build(
		stores.NewTodoStore,
		services.NewTodoService,
		handlers.NewTodoHandler,
	)

	return nil
}
