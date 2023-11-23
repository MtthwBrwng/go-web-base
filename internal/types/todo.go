package types

import "time"

type Todo struct {
	ID      string
	Created time.Time
	Updated time.Time
	Body    string
}

type NewTodoFormInput struct {
	Body string `form:"body"`
}

type EditTodoFormInput struct {
	Body string `form:"body"`
}
