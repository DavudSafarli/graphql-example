package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stdapps/graphql-example/delivery/graphql/graph/dto"
	"github.com/stdapps/graphql-example/delivery/graphql/graph/generated"
)

func (r *queryResolver) Users(ctx context.Context, pagination *dto.PaginationInput, criteria *dto.UsersCriteriaInput) ([]dto.User, error) {
	p := dto.MapToTicketingPagination(pagination)
	c := dto.MapToUserSearchCriteria(criteria)

	users, err := r.Storage.GetUsers(p, c)
	if err != nil {
		return nil, err
	}
	response := make([]dto.User, 0, len(users))
	for _, u := range users {
		response = append(response, dto.MapUser(u))
	}
	return response, err
}

func (r *queryResolver) User(ctx context.Context, id int) (*dto.User, error) {
	user, err := r.Storage.FindUser(id)
	if err != nil {
		return nil, err
	}
	response := dto.MapUser(user)
	return &response, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
