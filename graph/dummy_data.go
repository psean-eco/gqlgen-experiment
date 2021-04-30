package graph

import "github.com/psean/gql-go-gen/graph/model"

var home string = "My Home"
var devices = []*model.Device{
	{
		ID:       "1234",
		Name:     "test Device1",
		Location: &home,
	},
	{
		ID:       "5678",
		Name:     "test Device2",
		Location: &home,
	},
	{
		ID:       "0123",
		Name:     "test Device3",
		Location: &home,
	},
	{
		ID:       "65789",
		Name:     "test Device4",
		Location: &home,
	},
}
