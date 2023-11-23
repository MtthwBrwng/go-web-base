package stores

import (
	"context"
	"go-web-base/internal/types"
)

type TodoStore struct {
	store map[string]*types.Todo
}

func NewTodoStore(store map[string]*types.Todo) *TodoStore {
	return &TodoStore{store: store}
}

func (s *TodoStore) Create(ctx context.Context, input types.Todo) (*types.Todo, error) {
	s.store[input.ID] = &input
	return &input, nil
}

func (s *TodoStore) Update(ctx context.Context, id string, input types.Todo) (*types.Todo, error) {
	s.store[id] = &input
	return &input, nil
}

func (s *TodoStore) Delete(ctx context.Context, id string) error {
	delete(s.store, id)

	return nil
}

func (s *TodoStore) GetByID(ctx context.Context, id string) (*types.Todo, error) {
	return s.store[id], nil
}

func (s *TodoStore) GetAll(ctx context.Context) ([]*types.Todo, error) {
	var all []*types.Todo
	for _, todo := range s.store {
		all = append(all, todo)
	}

	return all, nil
}
