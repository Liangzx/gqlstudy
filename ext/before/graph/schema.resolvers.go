package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ooxx.com/ext/before/graph/generated"
	"ooxx.com/ext/before/graph/model"
)

func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "item-id-01", nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "user-id-01", nil
}

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	fmt.Printf("id: %s\n", id)
	return &model.Item{
		ID:   "item-id-01",
		Name: "item01",
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	fmt.Printf("id: %s\n", id)
	return &model.User{
		ID:   "user-id-01",
		Name: "user01",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
