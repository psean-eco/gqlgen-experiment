package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/psean/gql-go-gen/graph/generated"
	"github.com/psean/gql-go-gen/graph/model"
)

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	var device model.Device
	device.ID = input.ID
	device.Name = input.Name
	devices = append(devices, &device)
	return &device, nil
}

func (r *mutationResolver) DeleteDevice(ctx context.Context, input model.NewDevice) (bool, error) {
	for i, device := range devices {
		if input.ID == device.ID {
			devices = append(devices[:i], devices[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	return devices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
