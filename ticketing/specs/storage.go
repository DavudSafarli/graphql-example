package specs

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stdapps/graphql-example/ticketing"
	"github.com/stretchr/testify/require"
)

// TestStorage ...
func TestStorage(t *testing.T, storage ticketing.Storage) {
	name := fmt.Sprintf("randName-%d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(500))
	log.Println(name)
	user := ticketing.User{
		Name: name,
	}
	user, err := storage.CreateUser(user)
	require.Nil(t, err)

	retrieved, err := storage.FindUser(user.ID)
	require.Nil(t, err)

	require.Equal(t, retrieved, user)
}
