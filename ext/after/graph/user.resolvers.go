package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ooxx.com/ext/after/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "user-id-01", nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	fmt.Printf("id: %s\n", id)
	return &model.User{
		ID:   "user-id-01",
		Name: "user01",
	}, nil
}
