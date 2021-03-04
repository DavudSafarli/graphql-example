local-test-env-up:
	docker-compose -f ./docker-compose-test.yml up -d
	migrate -source "file://./storage/migrations" -database "postgres://gqluser:gqlpass@localhost:5433/gqltest?sslmode=disable" up

local-test-env-down:
	docker-compose -f ./docker-compose-test.yml down

local-test:
	make local-test-env-up
	go test ./... -p=1 --cover
