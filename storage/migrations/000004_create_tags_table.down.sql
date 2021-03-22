-- migrate -source file://./storage/migrations/ -database postgres://gqluser:gqlpass@localhost:5433/gqltest?sslmode=disable drop -f
DROP TABLE IF EXISTS tags;