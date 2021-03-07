package resolvers

import "github.com/stdapps/graphql-example/ticketing"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	Storage ticketing.Storage
}
