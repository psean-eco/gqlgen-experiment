package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/psean/gql-go-gen/graph/generated"
	"github.com/psean/gql-go-gen/graph/model"
	utils "github.com/psean/gql-go-gen/graph/utils"
)

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	res, err := utils.GetPlayers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
