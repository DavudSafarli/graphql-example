package dto

import "github.com/stdapps/graphql-example/ticketing"

// MapToTicketingPagination maps graphql dto to ticketing.Pagination
func MapToTicketingPagination(p *PaginationInput) ticketing.Pagination {
	if p == nil {
		return ticketing.Pagination{Limit: 10, Page: 0}
	}
	def := ticketing.Pagination{Limit: 10, Page: 0}
	if p.Limit != nil {
		def.Limit = *p.Limit
	}
	if p.Page != nil {
		def.Page = *p.Page
	}
	return def
}

// MapToUserSearchCriteria maps graphql dto to ticketing.ToUserSearchCriteria
func MapToUserSearchCriteria(c *UsersCriteriaInput) ticketing.UsersSearchCriteria {
	if c == nil {
		return ticketing.UsersSearchCriteria{}
	}
	def := ticketing.UsersSearchCriteria{}
	if c.Name != nil {
		def.Name = *c.Name
	}
	return def
}

// MapUser maps ticketing.User (domain model) to dto.User
func MapUser(u ticketing.User) User {
	return User{
		ID:   u.ID,
		Name: u.Name,
	}
}
