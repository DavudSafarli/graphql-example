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

// MapUser maps ticketing.User (domain model) to dto.User
func MapUsers(users []ticketing.User) []User {
	res := []User{}
	for _, v := range users {
		res = append(res, MapUser(v))
	}
	return res
}

// MapTicket maps ticketing.Ticket (domain model) to dto.Ticket
func MapTicket(u ticketing.Ticket) Ticket {
	return Ticket{
		ID:    u.ID,
		Title: u.Title,
	}
}

// MapTag maps ticketing.Tag (domain model) to dto.Tag
func MapTag(t ticketing.Tag) Tag {
	return Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

// MapTags maps ticketing.Tag (domain model) to dto.Tag
func MapTags(tags []ticketing.Tag) []Tag {
	res := []Tag{}
	for _, v := range tags {
		res = append(res, MapTag(v))
	}
	return res
}
