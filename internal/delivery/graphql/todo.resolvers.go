package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ent_example/internal/delivery/graphql/generated"
	"ent_example/internal/delivery/graphql/model"
	"ent_example/internal/ent"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*ent.Todo, error) {
//ent.TodoMutation.SetStatus()
	todoCreated, err := r.client.Todo.Create().SetStatus(todo.Status).SetText(todo.Text).SetPriority(*todo.Priority).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &ent.Todo{
		ID:        todoCreated.ID,
		Text:      todoCreated.Text,
		CreatedAt: todoCreated.CreatedAt,
		Status:    todoCreated.Status,
		Priority:  todoCreated.Priority,
	}, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	todos, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return todos, err
}

func (r *todoResolver) Parent(ctx context.Context, obj *ent.Todo) (*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo) ([]*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
