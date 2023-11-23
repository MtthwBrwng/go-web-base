package routes

import "go-web-base/internal/handlers"

type Service struct {
	TodoHandler *handlers.TodoHandler
}

func NewService(s Service) *Service {
	return &s
}
