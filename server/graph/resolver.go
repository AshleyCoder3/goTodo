package graph

import (
	"context"
	"github.com/ashleycoder3/go-react-todo/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

func (r *queryResolver) Todo(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented")
}
