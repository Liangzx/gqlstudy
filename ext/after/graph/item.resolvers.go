package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ooxx.com/ext/after/graph/model"
)

func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "item-id-01", nil
}

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	fmt.Printf("id: %s\n", id)
	return &model.Item{
		ID:   "item-id-01",
		Name: "item01",
	}, nil
}
