package specs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stdapps/graphql-example/ticketing"
	"github.com/stretchr/testify/require"
)

// TestStorage ...
func TestStorage(t *testing.T, storage ticketing.Storage) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	getRandomStr := func() string { return fmt.Sprintf("randName-%d", r.Int()) }

	t.Run(`#CreateUser created user should be findable with #FindUser`, func(t *testing.T) {
		user := ticketing.User{
			Name: getRandomStr(),
		}
		user, err := storage.CreateUser(user)
		require.Nil(t, err)

		retrieved, err := storage.FindUser(user.ID)
		require.Nil(t, err)

		require.Equal(t, retrieved, user)
	})

	t.Run(`#FindUser should return error if user doesn't exist`, func(t *testing.T) {
		impossibleToExist := -1
		_, err := storage.FindUser(impossibleToExist)
		require.NotNil(t, err)
	})

	t.Run(`#GetUsers should return non-empty slice for existing N records`, func(t *testing.T) {
		srchName := fmt.Sprint(r.Intn(99999))
		storage.CreateUser(ticketing.User{Name: "txt1" + srchName})
		storage.CreateUser(ticketing.User{Name: srchName + "txt2"})

		p := ticketing.Pagination{Limit: 3}
		l := ticketing.UsersSearchCriteria{Name: srchName}

		users, err := storage.GetUsers(p, l)
		require.Nil(t, err)
		require.Equal(t, 2, len(users))
	})

}
