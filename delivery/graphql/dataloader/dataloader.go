package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/stdapps/graphql-example/delivery/graphql/graph/dto"
	"github.com/stdapps/graphql-example/ticketing"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type DataLoader struct {
	storage ticketing.Storage
}

func NewDataLoader(storage ticketing.Storage) *DataLoader {
	return &DataLoader{storage: storage}
}

type loaders struct {
	AssigneesLoder *AssigneesLoader
	TagsLoader     *TagsLoader
}

func (dl *DataLoader) LoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		// set this to zero what happens without dataloading
		wait := 1 * time.Microsecond

		// M:M loader
		ldrs.AssigneesLoder = &AssigneesLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([][]dto.User, []error) {
				var ticketIds []int
				for _, key := range keys {
					ticketIds = append(ticketIds, key)
				}
				ticketAssignees, err := dl.storage.GetTicketsAssignees(ticketIds)
				if err != nil {
					panic(err)
				}

				items := make([][]dto.User, len(keys))
				errors := make([]error, len(keys))
				for i, tId := range keys {
					items[i] = dto.MapUsers(ticketAssignees[tId])
				}

				return items, errors
			},
		}

		ldrs.TagsLoader = &TagsLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([][]dto.Tag, []error) {
				var ticketIds []int
				for _, key := range keys {
					ticketIds = append(ticketIds, key)
				}
				ticketTags, err := dl.storage.GetTicketsTags(ticketIds)
				if err != nil {
					panic(err)
				}

				items := make([][]dto.Tag, len(keys))
				errors := make([]error, len(keys))
				for i, tId := range keys {
					items[i] = dto.MapTags(ticketTags[tId])
				}

				return items, errors
			},
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func For(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
