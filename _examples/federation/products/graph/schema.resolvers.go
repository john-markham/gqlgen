package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/john-markham/gqlgen version v0.17.64-dev

import (
	"context"

	"github.com/john-markham/gqlgen/_examples/federation/products/graph/model"
)

// TopProducts is the resolver for the topProducts field.
func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*model.Product, error) {
	return hats, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
