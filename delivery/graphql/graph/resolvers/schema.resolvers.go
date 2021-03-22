package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stdapps/graphql-example/delivery/graphql/dataloader"
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

func (r *queryResolver) Tickets(ctx context.Context, pagination *dto.PaginationInput) ([]dto.Ticket, error) {
	p := dto.MapToTicketingPagination(pagination)
	tickets, err := r.Storage.GetTickets(p)
	if err != nil {
		return nil, err
	}
	response := make([]dto.Ticket, 0, len(tickets))
	for _, t := range tickets {
		response = append(response, dto.MapTicket(t))
	}
	return response, err
}

func (r *queryResolver) Tags(ctx context.Context) ([]dto.Tag, error) {
	tags, err := r.Storage.GetTags()
	if err != nil {
		return nil, err
	}
	response := make([]dto.Tag, 0, len(tags))
	for _, t := range tags {
		response = append(response, dto.MapTag(t))
	}
	return response, err
}

func (r *ticketResolver) Assignees(ctx context.Context, obj *dto.Ticket) ([]dto.User, error) {
	return dataloader.For(ctx).AssigneesLoder.Load(obj.ID)
}

func (r *ticketResolver) Tags(ctx context.Context, obj *dto.Ticket) ([]dto.Tag, error) {
	return dataloader.For(ctx).TagsLoader.Load(obj.ID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Ticket returns generated.TicketResolver implementation.
func (r *Resolver) Ticket() generated.TicketResolver { return &ticketResolver{r} }

type queryResolver struct{ *Resolver }
type ticketResolver struct{ *Resolver }
