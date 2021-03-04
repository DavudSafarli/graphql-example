package storage

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stdapps/graphql-example/ticketing/specs"
	"github.com/stretchr/testify/require"
)

func TestPostgresStorage(t *testing.T) {
	connstr := "postgres://gqluser:gqlpass@localhost:5433/gqltest?sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	require.Nil(t, err)
	postgresStorage := NewPostgresStorage(db)

	specs.TestStorage(t, postgresStorage)

}
