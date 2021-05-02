package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/psean/gql-go-gen/graph/generated"
	"github.com/psean/gql-go-gen/graph/model"
	utils "github.com/psean/gql-go-gen/graph/utils"
)

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	res, err := utils.GetTeams()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	res, err := utils.GetPlayers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *teamResolver) Players(ctx context.Context, obj *model.Team) ([]*model.Player, error) {
	var players []*model.Player
	players, err := utils.GetPlayersForTeam(*obj)
	if err != nil {
		return players, err
	}
	return players, nil

}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type queryResolver struct{ *Resolver }
type teamResolver struct{ *Resolver }
